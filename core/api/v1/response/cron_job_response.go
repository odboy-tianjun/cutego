package response

type CronJobResponse struct {
	JobId     int    `json:"jobId"`     // 任务主键
	JobName   string `json:"jobName"`   // 任务名称
	JobCron   string `json:"jobCron"`   // cron表达式
	FuncAlias string `json:"funcAlias"` // 方法别名(程序内注册的别名)
	FuncParam string `json:"funcParam"` // 方法参数
	Status    string `json:"status"`    // 状态(1、Running 0、Stop)
	Level     int    `json:"level"`     // 任务级别(0、普通  1、一般 2、重要 3、强保)
	Remark    string `json:"remark"`    // 备注
}
