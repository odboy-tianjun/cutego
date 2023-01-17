package main

// init函数执行顺序自上而下, 最后执行main包里面的init函数
import (
	_ "cutego/core/dao"
	"cutego/core/router"
	"cutego/pkg/common"
	"cutego/pkg/config"
	_ "cutego/pkg/cronjob"
	"cutego/pkg/middleware/logger"
	"cutego/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	StartTest()
	StartApp()
}
func StartTest() {
	fmt.Println("================ Test Content =================")

	//cronjob.PrintCronNext()
	//cronjob.ExecWithCronNext()

	fmt.Println("================ Test Content =================")
}

func StartApp() {
	//switch config.AppEnvConfig.Server.RunMode {
	//case gin.DebugMode:
	//	gin.SetMode(gin.DebugMode)
	//	break
	//case gin.ReleaseMode:
	//	gin.SetMode(gin.ReleaseMode)
	//	break
	//default:
	//	gin.SetMode(gin.DebugMode)
	//}
	gin.SetMode(util.IF(config.AppEnvConfig.Server.RunMode == "", "debug", config.AppEnvConfig.Server.RunMode).(string))
	r := router.Init()
	r.Use(logger.LoggerToFile())
	err := r.Run(fmt.Sprintf(":%d", config.AppEnvConfig.Server.Port))
	if err != nil {
		common.FatalfLog("Start server: %+v", err)
	}
}
