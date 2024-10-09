package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	//to show the original place that calls the log
	// log, err= zap.NewProduction(zap.AddCallerSkip(1))

	//set the log content format
	config:=zap.NewProductionConfig()
	encoderConfig:=zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey="timestamp"
	encoderConfig.EncodeTime=zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey=""
	config.EncoderConfig=encoderConfig

	log, err=config.Build(zap.AddCallerSkip(1))

	if err !=nil {
		panic(err)
	}
}

func Info(msg string, fields ...zap.Field){
		log.Info(msg)
}

func Debug(msg string, fields ...zap.Field){
	log.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field){
	log.Error(msg, fields...)
}


//{"level":"info","ts":1728446488.5675535,"caller":"FirstGoServer/main.go:13","msg":"Starting the app..."}
//{"level":"info","timestamp":"2024-10-09T20:14:21.595+1300","caller":"FirstGoServer/main.go:13","msg":"Starting the app..."}