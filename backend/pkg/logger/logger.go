package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.SugaredLogger
var currentDate string
var logFileBase string

// Init 初始化 logger
// logFile: 日志文件路径（不带日期）
func Init(logFile string) {
	logFileBase = logFile
	currentDate = time.Now().Format("2006-01-02")
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileWithDate(),
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     30, // days
		Compress:   true,
	})

	// 时间格式统一：YYYY-MM-dd hh:mm:ss
	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	// 文件 JSON 编码
	fileEncoderConfig := zap.NewProductionEncoderConfig()
	fileEncoderConfig.TimeKey = "time"
	fileEncoderConfig.EncodeTime = timeEncoder
	fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(fileEncoderConfig),
		fileWriter,
		zap.InfoLevel,
	)

	// 控制台 彩色编码
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderConfig.EncodeTime = timeEncoder
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)

	// 合并 core
	core := zapcore.NewTee(fileCore, consoleCore)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log = logger.Sugar()
}

// 获取带日期的日志文件名
func logFileWithDate() string {
	return logFileBase + "-" + currentDate + ".log"
}

// 检查日期是否变化，变化则重新初始化 logger
func checkDateAndRotate() {
	today := time.Now().Format("2006-01-02")
	if today != currentDate {
		Init(logFileBase)
	}
}

// 对外暴露的方法
func Info(args ...interface{}) {
	checkDateAndRotate()
	log.Info(args...)
}
func Infof(template string, args ...interface{}) {
	checkDateAndRotate()
	log.Infof(template, args...)
}
func Debug(args ...interface{}) {
	checkDateAndRotate()
	log.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	checkDateAndRotate()
	log.Debugf(template, args...)
}
func Error(args ...interface{}) {
	checkDateAndRotate()
	log.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	checkDateAndRotate()
	log.Errorf(template, args...)
}
func Fatal(args ...interface{}) {
	checkDateAndRotate()
	log.Fatal(args...)
}
func Fatalf(template string, args ...interface{}) {
	checkDateAndRotate()
	log.Fatalf(template, args...)
}
