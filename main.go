package main

// init函数执行顺序自上而下, 最后执行main包里面的init函数
import (
	"cutego/pkg/config"
	_ "cutego/pkg/cronjob"
	_ "cutego/pkg/gin"
	"cutego/pkg/logging"
	"cutego/pkg/util"
	"cutego/refs"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//go testChangeJob()
	gin.SetMode(util.IF(config.AppEnvConfig.Server.RunMode == "", "debug", config.AppEnvConfig.Server.RunMode).(string))
	err := refs.CoolGin.Run(fmt.Sprintf(":%d", config.AppEnvConfig.Server.Port))
	if err != nil {
		logging.FatalfLog("Start server: %+v", err)
	}
}

//func testChangeJob() {
//	time.Sleep(time.Millisecond * 5000)
//	fmt.Println("改变任务调度间隔")
//	cronjob.AppendCronFunc("*/5 * * * *", "test1", "1")
//}
