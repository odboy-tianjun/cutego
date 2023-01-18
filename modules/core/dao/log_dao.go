package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/entity"
	"cutego/pkg/common"
	"cutego/pkg/page"
	"cutego/refs"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
)

type LogDao struct {
}

func (d LogDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_log")
}

// SelectPage 分页查询数据
func (d LogDao) SelectPage(query request.LogQuery) ([]entity.SysLog, int64) {
	configs := make([]entity.SysLog, 0)
	session := d.sql(refs.SqlDB.NewSession())
	if gotool.StrUtils.HasNotEmpty(query.Content) {
		session.And("content like concat('%', ?, '%')", query.Content)
	}
	if gotool.StrUtils.HasNotEmpty(query.Uid) {
		session.And("uid = ?", query.Uid)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&configs)
	if err != nil {
		common.ErrorLog(err)
		return nil, 0
	}
	return configs, total
}

// Insert 添加数据
func (d LogDao) Insert(config entity.SysLog) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&config)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}
