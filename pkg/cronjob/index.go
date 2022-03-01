package cronjob

import (
	"cutego/core/api/v1/request"
	"cutego/core/job"
	"cutego/core/service"
	"cutego/pkg/common"
	"github.com/robfig/cron"
	"time"
)

// Cron表达式参考
// 每隔5秒执行一次：*/5 * * * * ?
// 每隔1分钟执行一次：0 */1 * * * ?
// 每天23点执行一次：0 0 23 * * ?
// 每天凌晨1点执行一次：0 0 1 * * ?
// 每月1号凌晨1点执行一次：0 0 1 1 * ?
// 每月最后一天23点执行一次：0 0 23 L * ?
// 每周星期天凌晨1点实行一次：0 0 1 ? * L
// 在26分、29分、33分执行一次：0 26,29,33 * * * ?
// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?

// 定时任务: 别名与调度器的映射
var AliasCronMap = make(map[string]*cron.Cron)

// 停止任务, 不会停止已开始的任务
func StopCronFunc(aliasName string) {
	common.InfoLogf("停止任务 %s ---> Start", aliasName)
	AliasCronMap[aliasName].Stop()
	common.InfoLogf("停止任务 %s ---> Finish", aliasName)
}

// 开始任务
func StartCronFunc(aliasName string) {
	common.InfoLogf("唤起任务 %s ---> Start", aliasName)
	AliasCronMap[aliasName].Start()
	common.InfoLogf("唤起任务 %s ---> Finish", aliasName)
}

func init() {
	if len(job.AliasFuncMap) > 0 {
		//go test()
		index := 1
		for true {
			q := request.CronJobQuery{}
			q.PageNum = index
			data, _ := service.CronJobService{}.FindPage(q)
			if len(data) == 0 {
				break
			}
			for _, datum := range data {
				c := cron.New()
				c.AddFunc(datum.JobCron, job.AliasFuncMap[datum.FuncAlias])
				c.Start()

				AliasCronMap[datum.FuncAlias] = c
				common.InfoLogf("调度定时任务 --- %s ---> Success", datum.JobName)
			}
			index += 1
		}
	}
}

// 测试通过
func test() {
	time.Sleep(time.Second * 10)
	StopCronFunc("test1")
	time.Sleep(time.Second * 10)
	StartCronFunc("test1")
}
