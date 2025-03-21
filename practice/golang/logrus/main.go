package main

import (
	"bytes"
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

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

func init() {
	logrus.SetFormatter(&BaseFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.WarnLevel)

	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.WithField("error", err).Error("Failed to log to file")
	}

	writers := io.MultiWriter(file,
		os.Stdout, &lumberjack.Logger{
			Filename:   "./log/test.log",
			MaxSize:    1,
			MaxBackups: 3,
			MaxAge:     1,
			Compress:   false,
		})

	logrus.SetOutput(writers)
}

func main() {
	logrus.WithFields(logrus.Fields{
		"ip": "127.0.0.1",
	}).Info("info log")

	for i := 0; i < 2; i++ {
		logrus.WithFields(logrus.Fields{
			"ip":    "127.0.0.1",
			"index": i,
		}).Error("error log")
	}
}
