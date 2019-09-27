package logdef

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"time"
)

var Logger *zap.Logger

type LogInfo struct {
	LogPath  string `json:"log_path" toml:"log_path"`
	LogLevel string `json:"log_level" toml:"log_level"`
	LogAddr  string `json:"log_addr" toml:"log_addr"`
}

func HumanFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
func logLevelListen(handler func(http.ResponseWriter, *http.Request), addr string) {
	logLevelHandle := http.NewServeMux()
	logLevelHandle.HandleFunc("/handle/loglevel", handler)
	err := http.ListenAndServe(addr, logLevelHandle)
	if err != nil {
		panic(err)
	}
}
func Init(loginf *LogInfo) *zap.Logger {
	hook := lumberjack.Logger{Filename: loginf.LogPath, MaxSize: 500, MaxBackups: 5, MaxAge: 3, LocalTime: true, Compress: false}
	w := zapcore.AddSync(&hook)
	var level zap.AtomicLevel
	switch loginf.LogLevel {
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		level = zap.NewAtomicLevel()
	}
	go logLevelListen(level.ServeHTTP, loginf.LogAddr)
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = HumanFormat
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), w, level)
	Logger = zap.New(core, zap.AddCaller())
	return Logger
}
