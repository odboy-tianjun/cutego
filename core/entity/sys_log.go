package entity

import (
	"time"
)

type SysLog struct {
	Id         int64     `xorm:"pk autoincr" json:"id"`       // 主键
	Uid        string    `xorm:"uid" json:"uid"`              // 唯一主键
	Content    string    `xorm:"varchar(255)" json:"content"` // 记录内容
	CreateTime time.Time `xorm:"created"`                     // 创建时间
}

func (c SysLog) TableName() string {
	return "sys_log"
}
