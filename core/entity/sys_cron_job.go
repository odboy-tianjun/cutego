package entity

import (
	"time"
)

type SysCronJob struct {
	JobId      int64     `xorm:"pk autoincr" json:"jobId"`      // 任务主键
	JobName    string    `xorm:"varchar(100)" json:"jobName"`   // 任务名称
	JobCron    string    `xorm:"varchar(255)" json:"jobCron"`   // cron表达式
	FuncAlias  string    `xorm:"varchar(100)" json:"funcAlias"` // 方法别名(程序内注册的别名)
	Status     string    `xorm:"char(1)" json:"status"`         // 状态(1、Running 0、Stop)
	Level      int       `xorm:"int(1)" json:"level"`           // 任务级别(0、普通  1、一般 2、重要 3、强保)
	CreateBy   string    `xorm:"varchar(64)"`                   // 创建人
	CreateTime time.Time `xorm:"created"`                       // 创建时间
	UpdateBy   string    `xorm:"varchar(64)"`                   // 更新人
	UpdateTime time.Time `xorm:"datetime"`                      // 更新时间
	Remark     string    `xorm:"varchar(500)" json:"remark"`    // 备注
}

func (c SysCronJob) TableName() string {
	return "sys_cron_job"
}
