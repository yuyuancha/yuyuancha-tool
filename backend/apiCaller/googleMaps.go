package apiCaller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yuyuancha/yuyuancha-tool/config"
	"io"
	"net/http"
	"strings"
	"time"
)

// GoogleMapsApiCaller google maps api 接口
type GoogleMapsApiCaller struct {
	client *http.Client
}

// googleMapsApi google maps api
type googleMapsApi struct {
	method string
	url    string
}

// googleMapsBaseResponse google maps 基礎回應格式
type googleMapsBaseResponse struct {
	Results interface{} `json:"results"`
	Status  string      `json:"status"`
}

// GoogleMapsGetAddressInfoResponse 取得地址資訊回應
type GoogleMapsGetAddressInfoResponse struct {
	Geometry struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
}

var googleMapsApis = map[string]googleMapsApi{
	"getAddressInfo": {http.MethodGet, "https://maps.googleapis.com/maps/api/geocode/json"},
}

// NewGoogleMapsApiCaller 建立 google maps api
func NewGoogleMapsApiCaller() *GoogleMapsApiCaller {
	client := &http.Client{}
	client.Timeout = time.Duration(HttpTimeout) * time.Second

	return &GoogleMapsApiCaller{
		client: client,
	}
}

// GetAddressLatAndLng 取得地址緯度和經度
func (apiCaller *GoogleMapsApiCaller) GetAddressLatAndLng(address string) (float64, float64) {
	var params = map[string]string{
		"address": address,
		"key":     config.GoogleMapsApiKey,
	}

	var results []GoogleMapsGetAddressInfoResponse

	_, err := apiCaller.doRequest(googleMapsApis["getAddressInfo"], params, &results)
	if err != nil {
		fmt.Println("Google maps api get address info error:", err.Error())
		return 0, 0
	}

	if len(results) < 1 {
		fmt.Println("Google maps api get address info results is empty.")
		return 0, 0
	}

	return results[0].Geometry.Location.Lat, results[0].Geometry.Location.Lng
}

// doRequest 呼叫請求
func (apiCaller *GoogleMapsApiCaller) doRequest(api googleMapsApi, params map[string]string, result interface{}) (responseInfo googleMapsBaseResponse, err error) {
	var requestUrl = api.url
	var body io.Reader = nil

	switch api.method {
	case http.MethodGet:
		requestUrl += apiCaller.combineGetRequestUrl(params)
	default:
		fmt.Println("Api action is not existed.")
		return responseInfo, errors.New("api action is not existed")
	}

	request, err := http.NewRequest(api.method, requestUrl, body)
	if err != nil {
		fmt.Println("Google maps api request error:", err.Error())
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err := apiCaller.client.Do(request)
	if err != nil {
		fmt.Println("Google maps api client do error:", err.Error())
		return
	}

	defer func() {
		closeErr := response.Body.Close()
		if closeErr != nil {
			fmt.Println("Google maps api response body close err:", closeErr)
		}
	}()

	responseInfo, err = apiCaller.parseResponse(response, result)
	if err != nil {
		fmt.Println("Parse Google maps api response error:", err.Error())
	}

	return
}

// combineGetRequestUrl 組成 GET 請求 url
func (apiCaller *GoogleMapsApiCaller) combineGetRequestUrl(params map[string]string) string {
	var results []string

	for key, value := range params {
		results = append(results, key+"="+value)
	}

	return "?" + strings.Join(results, "&")
}

// parseResponse 解析 http 回復信息
func (apiCaller *GoogleMapsApiCaller) parseResponse(response *http.Response, result interface{}) (googleMapsBaseResponse, error) {
	var responseInfo googleMapsBaseResponse

	if response.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("Google maps api http code: (%d)", response.StatusCode)
		fmt.Println(errorMessage)
		return responseInfo, errors.New(errorMessage)
	}

	if err := json.NewDecoder(response.Body).Decode(&responseInfo); err != nil {
		fmt.Println("解析 Google maps api 回復信息 Body Decode 錯誤:", err.Error())
		return responseInfo, err
	}

	responseInfoBytes, _ := json.Marshal(responseInfo)
	fmt.Println("GoogleMapsApiResponseInfo:", string(responseInfoBytes))

	responseDataBytes, err := json.Marshal(responseInfo.Results)
	if err != nil {
		return responseInfo, err
	}

	err = json.Unmarshal(responseDataBytes, result)
	if err != nil {
		return responseInfo, errors.New("google maps api response data unmarshal error")
	}

	return responseInfo, err
}
