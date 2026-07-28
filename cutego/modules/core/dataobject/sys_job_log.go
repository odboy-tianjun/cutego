package dataobject

import "time"

type SysJobLog struct {
	JobLogId     int64     `xorm:"pk autoincr" json:"jobLogId"`     // 日志主键
	JobId        int64     `xorm:"int(20)" json:"jobId"`            // 任务ID
	JobName      string    `xorm:"varchar(100)" json:"jobName"`     // 任务名称
	FuncAlias    string    `xorm:"varchar(100)" json:"funcAlias"`   // 调用目标
	JobMessage   string    `xorm:"varchar(500)" json:"jobMessage"`  // 日志信息
	Status       string    `xorm:"char(1)" json:"status"`           // 执行状态(0正常 1失败)
	ExceptionInfo string   `xorm:"varchar(2000)" json:"exceptionInfo"` // 异常信息
	CreateTime   time.Time `xorm:"created" json:"createTime"`       // 创建时间
}

func (s SysJobLog) TableName() string {
	return "sys_job_log"
}