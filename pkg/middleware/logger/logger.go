package logger

import (
	"cutego/pkg/config"
	"cutego/pkg/logging"
	"cutego/pkg/util"
	"fmt"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	dirPath := config.GetDirPath("logging")
	fileName := path.Join(dirPath, "application.logging")
	if !util.IsFileOrDirExist(dirPath) {
		err := util.CreateAllDir(dirPath)
		if err != nil {
			logging.ErrorLog(err)
		}
	}
	if !gotool.FileUtils.Exists(fileName) {
		create, err := os.Create(fileName)
		if err != nil {
			logging.ErrorLog(err)
		}
		defer create.Close()
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	// 输出源
	logger.Out = src
	switch config.AppEnvConfig.Logger.Level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
		break
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
		break
	case "info":
		logger.SetLevel(logrus.InfoLevel)
		break
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
		break
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
		break
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	default:
		logger.SetLevel(logrus.DebugLevel)
	}
	logger.SetLevel(logrus.DebugLevel)
	// 设置 rotatelogs
	if config.AppEnvConfig.Logger.MaxSaveAge == 0 {
		config.AppEnvConfig.Logger.MaxSaveAge = 7
	}
	if config.AppEnvConfig.Logger.RotationTime == 0 {
		config.AppEnvConfig.Logger.RotationTime = 1
	}
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.logs",
		// 生成软链, 指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Duration(config.AppEnvConfig.Logger.MaxSaveAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(config.AppEnvConfig.Logger.RotationTime)*24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增 Hook
	logger.AddHook(lfHook)
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
