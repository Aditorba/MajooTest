package log

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap/zapcore"
	"time"
)
import "go.uber.org/zap"

var (
	logs    *zap.Logger
	dateNow = time.Now()
)

func init() {
	var err error

	config := zap.NewProductionConfig()
	dateNow := time.Now()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.CallerKey = "function"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = TimeFormat
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.MessageKey = "message"
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.FunctionKey = "function key"
	config.EncoderConfig = encoderConfig

	//write log
	config.OutputPaths = []string{
		fmt.Sprintf("log/data/majoo-%s.log", dateNow.Format("2006-01-02")),
	}

	logs, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

}

func iso3339CleanTime(time.Time) string {
	date := dateNow.Format("2006-01-02 15:04:05")
	return date
}

func TimeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(iso3339CleanTime(t))
}

func Info(message string, data interface{}) {
	resultjs, _ := json.Marshal(data)
	jsRaw := json.RawMessage(string(resultjs)[:])
	logs.Info(message, zap.Any("info", &jsRaw))
}

func Debug(message string, data interface{}) {
	resultjs, _ := json.Marshal(data)
	jsRaw := json.RawMessage(string(resultjs)[:])
	logs.Debug(message, zap.Any("debug", &jsRaw))
}

func Error(message string, data interface{}) {
	resultjs, _ := json.Marshal(data)
	jsRaw := json.RawMessage(string(resultjs)[:])
	logs.Error(message, zap.Any("error", &jsRaw))

}
