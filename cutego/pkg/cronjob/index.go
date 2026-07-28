package cronjob

import (
	"cutego/modules/core/job"
	"cutego/modules/core/service"
	"cutego/pkg/logger"
	"sync"

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
var aliasCronMap sync.Map

// getCron 获取别名对应的 cron 实例
func getCron(aliasName string) *cron.Cron {
	v, ok := aliasCronMap.Load(aliasName)
	if !ok {
		return nil
	}
	return v.(*cron.Cron)
}

// StopCronFunc 停止任务, 不会停止已开始的任务
func StopCronFunc(aliasName string) {
	logger.SugaredLogger.Infof("停止任务 %s ---> Start", aliasName)
	if c := getCron(aliasName); c != nil {
		go c.Stop()
	}
	logger.SugaredLogger.Infof("停止任务 %s ---> Finish", aliasName)
}

// StartCronFunc 开始任务
func StartCronFunc(aliasName string) {
	logger.SugaredLogger.Infof("唤起任务 %s ---> Start", aliasName)
	if c := getCron(aliasName); c != nil {
		go c.Start()
	}
	logger.SugaredLogger.Infof("唤起任务 %s ---> Finish", aliasName)
}

// RemoveCronFunc 移除任务
func RemoveCronFunc(aliasName string) {
	logger.SugaredLogger.Infof("移除任务 %s ---> Start", aliasName)
	if c := getCron(aliasName); c != nil {
		go c.Stop()
	}
	aliasCronMap.Delete(aliasName)
	logger.SugaredLogger.Infof("移除任务 %s ---> Finish", aliasName)
}

// AppendCronFunc 新增任务
func AppendCronFunc(jobCron string, aliasName string, status string) {
	if c := getCron(aliasName); c != nil {
		c.Stop()
		aliasCronMap.Delete(aliasName)
	}
	logger.SugaredLogger.Infof("新增任务 %s ---> Start", aliasName)
	c := cron.New()
	err := c.AddFunc(jobCron, job.AliasFuncMap[aliasName])
	if err != nil {
		panic("任务追加失败, " + err.Error())
	}
	if status == "1" {
		go func() {
			c.Start()
			aliasCronMap.Store(aliasName, c)
			logger.SugaredLogger.Infof("调度定时任务 --- %s ---> Success", aliasName)
		}()
	} else {
		aliasCronMap.Store(aliasName, c)
	}
	logger.SugaredLogger.Infof("新增任务 %s ---> Finish", aliasName)
}

func InitJob() {
	jobService := service.CronJobService{}
	jobs, total := jobService.FindAll()
	if len(job.AliasFuncMap) > 0 && total > 0 {
		for _, datum := range jobs {
			AppendCronFunc(datum.JobCron, datum.FuncAlias, datum.Status)
		}
	}
}
