package dataobject

import (
	"time"
)

type Student struct {
	Id         int64     `xorm:"pk autoincr not null bigint 'id'"`
	CreateBy   string    `xorm:"default 'NULL' varchar(255) 'create_by'"`
	CreateTime time.Time `xorm:"default 'NULL' datetime 'create_time'"`
	UpdateBy   string    `xorm:"default 'NULL' varchar(255) 'update_by'"`
	UpdateTime time.Time `xorm:"default 'NULL' datetime 'update_time'"`
	Name       string    `xorm:"default 'NULL' varchar(255) 'name'"`
	Sex        string    `xorm:"default 'NULL' varchar(255) 'sex'"`
}

func (s *Student) TableName() string {
	return "student"
}
