package model

// WeatherResponseInfo 回復信息
type WeatherResponseInfo struct {
	Success string      `json:"success"`
	Records interface{} `json:"records"`
}

type WeatherOneWeekRecords struct {
	Locations []WeatherOneWeeKLocations `json:"locations"`
}

type WeatherOneWeeKLocations struct {
	Location []WeatherOneWeekLocation `json:"location"`
}

type WeatherOneWeekLocation struct {
	LocationName   string           `json:"locationName"`
	WeatherElement []WeatherElement `json:"weatherElement"`
}

type WeatherElement struct {
	ElementName string               `json:"elementName"`
	Description string               `json:"description"`
	Time        []WeatherElementTime `json:"time"`
}

type WeatherElementTime struct {
	StartTime    string                `json:"startTime"`
	EndTime      string                `json:"endTime"`
	ElementValue []WeatherElementValue `json:"elementValue"`
}

type WeatherElementValue struct {
	Value    string `json:"value"`
	Measures string `json:"measures"`
}
