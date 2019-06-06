package logdef

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"time"
)

var logger *zap.Logger

func HumanFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func logLevelListen(handler func(http.ResponseWriter, *http.Request),addr string){
	logLevelHandle := http.NewServeMux()
	logLevelHandle.HandleFunc("/handle/loglevel",handler)
	err := http.ListenAndServe(addr,logLevelHandle)
	if err != nil{
		panic(err)
	}
}
func Init(logpath string ,loglevel string ,addr string) * zap.Logger {
	hook := lumberjack.Logger{Filename:logpath,MaxSize:500,MaxBackups:5,MaxAge:3,LocalTime:true,Compress:false}
	w := zapcore.AddSync(&hook)
	var level zap.AtomicLevel

	switch loglevel {
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		level = zap.NewAtomicLevel()
	}

	go logLevelListen(level.ServeHTTP,addr)
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = HumanFormat
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder),w,level)
	logger := zap.New(core,zap.AddCaller())
	return logger
}

func LogDebugS(key string,value string,desc ...string){
	if len(desc) > 0{
		logger.Debug(desc[0],zap.String(key,value))
	}else{
		logger.Debug("",zap.String(key,value))
	}
}
func LogDebugI(key string,value int64, desc ...string ){
	if len(desc) > 0{
		logger.Debug(desc[0],zap.Int64(key,value))
	}else{
		logger.Debug("",zap.Int64(key,value))
	}
}
func LogErrorS(key string,value string,desc ...string){
	if len(desc) > 0{
		logger.Error(desc[0],zap.String(key,value))
	}else{
		logger.Error("",zap.String(key,value))
	}
}
func LogErrorI(key string,value int64, desc ...string ){
	if len(desc) > 0{
		logger.Error(desc[0],zap.Int64(key,value))
	}else{
		logger.Error("",zap.Int64(key,value))
	}
}

func LogInfoS(key string,value string,desc ...string){
	if len(desc) > 0{
		logger.Info(desc[0],zap.String(key,value))
	}else{
		logger.Info("",zap.String(key,value))
	}
}

func LogInfoI(key string,value int64, desc ...string ){
	if len(desc) > 0{
		logger.Info(desc[0],zap.Int64(key,value))
	}else{
		logger.Info("",zap.Int64(key,value))
	}
}
