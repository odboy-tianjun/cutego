package cronjob

import (
	"cutego/modules/core/job"
	"cutego/modules/core/service"
	"cutego/pkg/common"
	"github.com/robfig/cron"
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

// aliasCronMap 定时任务: 别名与调度器的映射
var aliasCronMap = make(map[string]*cron.Cron)

// StopCronFunc 停止任务, 不会停止已开始的任务
func StopCronFunc(aliasName string) {
	common.InfoLogf("停止任务 %s ---> Start", aliasName)
	go aliasCronMap[aliasName].Stop()
	common.InfoLogf("停止任务 %s ---> Finish", aliasName)
}

// StartCronFunc 开始任务
func StartCronFunc(aliasName string) {
	common.InfoLogf("唤起任务 %s ---> Start", aliasName)
	go aliasCronMap[aliasName].Start()
	common.InfoLogf("唤起任务 %s ---> Finish", aliasName)
}

// RemoveCronFunc 移除任务
func RemoveCronFunc(aliasName string) {
	common.InfoLogf("移除任务 %s ---> Start", aliasName)
	go StopCronFunc(aliasName)
	delete(aliasCronMap, aliasName)
	common.InfoLogf("移除任务 %s ---> Finish", aliasName)
}

// AppendCronFunc 新增任务
func AppendCronFunc(jobCron string, aliasName string, status string) {
	if aliasCronMap[aliasName] != nil {
		aliasCronMap[aliasName].Stop()
		aliasCronMap[aliasName] = nil
	}
	common.InfoLogf("新增任务 %s ---> Start", aliasName)
	c := cron.New()
	err := c.AddFunc(jobCron, job.AliasFuncMap[aliasName])
	if err != nil {
		panic("任务追加失败, " + err.Error())
	}
	if status == "1" {
		go func() {
			c.Start()
			aliasCronMap[aliasName] = c
			common.InfoLogf("调度定时任务 --- %s ---> Success", aliasName)
		}()
	}
	common.InfoLogf("新增任务 %s ---> Finish", aliasName)
}

func init() {
	jobService := service.CronJobService{}
	jobs, total := jobService.FindAll()
	if len(job.AliasFuncMap) > 0 && total > 0 {
		for _, datum := range jobs {
			AppendCronFunc(datum.JobCron, datum.FuncAlias, datum.Status)
		}
	}
}
