package apiCaller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yuyuancha/yuyuancha-tool/config"
	model "github.com/yuyuancha/yuyuancha-tool/model/weather"
	"io"
	"net/http"
	"time"
)

const HttpTimeout = 5
const PrintLimitLen = 1000

// WeatherApiCaller 天氣 api 呼叫器
type WeatherApiCaller struct {
	client *http.Client
}

var weatherApis map[string]weatherApi

func init() {
	weatherApis = map[string]weatherApi{
		"GetFeatherOneWeek": {http.MethodGet, "/v1/rest/datastore/F-D0047-091"},
	}
}

// weatherApi 天氣 API
type weatherApi struct {
	Action string
	Url    string
}

// Init 初始化 WeatherLogic
func (apiCaller *WeatherApiCaller) Init() {
	apiCaller.client = &http.Client{}
	apiCaller.client.Timeout = time.Duration(HttpTimeout) * time.Second
}

// GetOneWeek 取得一週天氣資料
func (apiCaller *WeatherApiCaller) GetOneWeek() ([]model.WeatherOneWeekLocation, error) {
	var result model.WeatherOneWeekRecords
	_, err := apiCaller.doRequest("GetFeatherOneWeek", &result)
	if err != nil {
		return nil, err
	}

	if len(result.Locations) == 0 {
		return nil, errors.New("取得一週天氣資料失敗: Locations 為空")
	}

	return result.Locations[0].Location, nil
}

// doRequest 發送請求
func (apiCaller *WeatherApiCaller) doRequest(apiName string, result interface{}) (responseInfo model.WeatherResponseInfo, err error) {
	var body io.Reader = nil

	request, err := http.NewRequest(weatherApis[apiName].Action, apiCaller.getWeatherApiUrl(apiName), body)
	if err != nil {
		fmt.Println("發送天氣", apiName, "請求失敗:", err.Error())
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err := apiCaller.client.Do(request)
	if err != nil {
		fmt.Println("發送天氣", apiName, "請求失敗:", err.Error())
		return
	}

	defer func() {
		closeErr := response.Body.Close()
		if closeErr != nil {
			fmt.Println("關閉天氣請求失敗:", closeErr)
		}
	}()

	responseInfo, err = apiCaller.parseResponse(response, result)

	return
}

// getWeatherApiUrl 取得天氣 API Url
func (apiCaller *WeatherApiCaller) getWeatherApiUrl(apiName string) string {
	return config.WeatherApi.Url + weatherApis[apiName].Url + "?Authorization=" + config.WeatherApi.AuthCode
}

// parseResponse 解析 http 回復信息
func (apiCaller *WeatherApiCaller) parseResponse(response *http.Response, result interface{}) (model.WeatherResponseInfo, error) {
	var responseInfo model.WeatherResponseInfo

	if response.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("解析天氣請求回應錯誤 http code: (%d)", response.StatusCode)
		fmt.Println(errorMessage)
		return responseInfo, errors.New(errorMessage)
	}

	if err := json.NewDecoder(response.Body).Decode(&responseInfo); err != nil {
		fmt.Println("解析天氣回復信息 Body Decode 錯誤:", err.Error())
		return responseInfo, err
	}

	responseInfoBytes, _ := json.Marshal(responseInfo)
	if len(responseInfoBytes) > PrintLimitLen {
		fmt.Println("天氣回應:", string(responseInfoBytes[:PrintLimitLen]))
	} else {
		fmt.Println("天氣回應:", string(responseInfoBytes))
	}

	if responseInfo.Records == nil {
		return responseInfo, errors.New("天氣回復信息 Records 為空")
	}

	responseRecordsBytes, err := json.Marshal(responseInfo.Records)
	if err != nil {
		return responseInfo, err
	}

	err = json.Unmarshal(responseRecordsBytes, result)
	if err != nil {
		return responseInfo, errors.New(responseInfo.Records.(string))
	}

	return responseInfo, err
}
