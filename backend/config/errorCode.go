package config

const (
	// 通用錯誤代碼
	ErrorCodeFormatValid = 90100
	ErrorCodeBadRequest  = 90101
)

var ErrorCodeMessage = map[int]string{
	ErrorCodeFormatValid: "格式錯誤",
	ErrorCodeBadRequest:  "資料處理失敗",
}
