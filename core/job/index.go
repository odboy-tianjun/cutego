package job

import (
	"cutego/core/service"
	"cutego/pkg/common"
)

// 定时任务: 别名与方法的映射
var AliasFuncMap = make(map[string]func())

// 注册任务
func RegisterFunc(aliasName string, f func()) {
	currentJob := service.CronJobService{}.GetInfoByAlias(aliasName)
	AliasFuncMap[aliasName] = f
	common.InfoLogf("注册定时任务 --- %s ---> Success", currentJob.JobName)
}

// 注册方法
func init() {
	//RegisterFunc("test1", TestJob)
}
