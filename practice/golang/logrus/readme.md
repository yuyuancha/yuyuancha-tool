## Logrus 練習

### 目的

練習使用 `logrus` 套件，以達到更好的 log 輸出效果。

### 用法

#### 安裝套件

```shell
go get github.com/sirupsen/logrus
```

#### 初始化設定

設定 `logrus` 的 formatter，可以使用 `TextFormatter` 或 `JSONFormatter`。

- `TimestampFormat` 可以自定義設定時間格式。

```
logrus.SetFormatter(&logrus.JSONFormatter{
	TimestampFormat: "2006-01-02 15:04:05",
})
```

設定是否輸出紀錄 log 的程式碼位置。

```
logrus.SetReportCaller(true)
```

設定輸出的警告級別，日誌級別高於設定輸出級別時，才會記錄日誌。

```
logrus.SetLevel(logrus.WarnLevel)
```

設定輸出位置，可以輸出到檔案或是標準輸出(打印在 terminal)。

```
file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
	logrus.WithField("error", err).Error("Failed to log to file")
}
writers := []io.Writer{
	file,
	os.Stdout,
}
fileAndStdoutWriter := io.MultiWriter(writers...)
logrus.SetOutput(fileAndStdoutWriter)
```

設定輸出位置，並限制檔案容量，超過容量時會自動切割檔案。

- 使用套件 `gopkg.in/natefinch/lumberjack.v2`。

```shell
go get -u github.com/natefinch/lumberjack
```

設定 `io.Writer`。
- Filename: 檔案名稱。
- MaxSize: 檔案最大容量。
- MaxBackups: 最大備份數量。
- MaxAge: 最大保存天數。
- Compress: 是否壓縮。

```
logrus.SetOutput(&lumberjack.Logger{
	Filename:   "./log/test.log",
	MaxSize:    1,
	MaxBackups: 3,
	MaxAge:     1,
	Compress:   false,
})
```

#### 使用

可以透過 `WithFields` 設定 log 的欄位，並呼叫所對應的級別方法。

對應日誌級別：
- Panic：紀錄日誌後 `panic()`。
- Fatal：紀錄日誌後 `os.Exit(1)`。
- Error：紀錄錯誤日誌。
- Warn：紀錄警告日誌。
- Info：紀錄一般日誌。
- Debug：紀錄 debug 日誌。
- Trace：紀錄 trace 日誌。

```
logrus.WithFields(logrus.Fields{
    "ip": "127.0.0.1",
}).Info("info log")

logrus.WithFields(logrus.Fields{
    "ip": "127.0.0.1",
}).Error("error log")
```

Output:
```
{"ip":"127.0.0.1","level":"info","msg":"info log","time":"2025-03-21 10:56:59"}
{"ip":"127.0.0.1","level":"error","msg":"error log","time":"2025-03-21 10:56:59"}
```

#### 自定義 log formatter

透過實作 `Formatter` 介面，可以自定義 log 的格式。

```
// BaseFormatter 基本格式
type BaseFormatter struct{}

// Format 實作格式化
func (formatter *BaseFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var (
		newLog     string
		timestamp  = entry.Time.Format("2006-01-02 15:04:05")
		byteBuffer = &bytes.Buffer{}
	)
	if entry.Buffer != nil {
		byteBuffer = entry.Buffer
	}

	if entry.HasCaller() {
		fileName := path.Base(entry.Caller.File)
		newLog = fmt.Sprintf("%s [%s] [%s:%d] %s\n",
			timestamp, entry.Level.String(), fileName, entry.Caller.Line, entry.Message)
	} else {
		newLog = fmt.Sprintf("%s [%s] %s\n", timestamp, entry.Level.String(), entry.Message)
	}

	byteBuffer.WriteString(newLog)
	return byteBuffer.Bytes(), nil
}
```

使用自定義的 formatter。

```
logrus.SetFormatter(&BaseFormatter{})
```

Output:
```
2025-03-21 15:11:24 [error] [main.go:70] error log
2025-03-21 15:11:24 [error] [main.go:70] error log
```
