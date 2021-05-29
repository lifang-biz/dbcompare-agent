package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"lifang.biz/dbcompare-client/conf"
)

var Logger *zap.Logger

func SetupLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()

	logLevel := zapcore.InfoLevel

	if conf.Config().App.RunMode == "test" {
		logLevel = zapcore.DebugLevel
	}else if conf.Config().App.RunMode == "pro" {
		logLevel = zapcore.ErrorLevel
	}

	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	Logger = zap.New(core, zap.AddCaller())
	Logger.Info( "zap Logger inited")
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./runtime/logs/latest.log",  // 日志输出文件
		MaxSize:    20,  // 日志最大保存1M
		MaxBackups: 50,  // 就日志保留50个日志备份
		MaxAge:     50,  // 最多保留50天日志
		Compress:   false, // 自导打 gzip包 默认false
	}
	return zapcore.AddSync(lumberJackLogger)
}