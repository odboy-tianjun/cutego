package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 全局 Logger
var globalLogger *zap.Logger
var SugaredLogger *zap.SugaredLogger

// InitZapLogger 初始化 Zap 生产级别日志
func InitZapLogger() {
	// 日志编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 彩色级别
		EncodeTime:     timeEncoder,                      // 时间格式化
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 控制台输出
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(
		consoleEncoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 输出到控制台
		zap.InfoLevel, // 日志级别
	)

	// 构建 logger
	globalLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	SugaredLogger = globalLogger.Sugar()

	// 替换全局 logger
	zap.ReplaceGlobals(globalLogger)
}

// 时间格式化
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// Info 对外方法
func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}

// Error 对外方法
func Error(msg string, fields ...zap.Field) {
	globalLogger.Error(msg, fields...)
}

// Debug 对外方法
func Debug(msg string, fields ...zap.Field) {
	globalLogger.Debug(msg, fields...)
}

// Warn 对外方法
func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}
