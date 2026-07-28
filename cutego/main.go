package main

// init函数执行顺序自上而下, 最后执行main包里面的init函数
import (
	"cutego/modules/core/dao"
	"cutego/pkg/config"
	"cutego/pkg/cronjob"
	_ "cutego/pkg/cronjob"
	"cutego/pkg/gin"
	_ "cutego/pkg/gin"
	"cutego/pkg/logger"
	"cutego/refs"
)

func main() {
	logger.InitZapLogger()
	config.InitConfig()
	refs.InitDatabase()
	refs.InitRedis()
	dao.PreInitConfig()
	dao.PreInitDictData()
	cronjob.InitJob()
	gin.InitServer()
}

//func testChangeJob() {
//	time.Sleep(time.Millisecond * 5000)
//	fmt.Println("改变任务调度间隔")
//	cronjob.AppendCronFunc("*/5 * * * *", "test1", "1")
//}
