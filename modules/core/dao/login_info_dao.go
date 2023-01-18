package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/pkg/common"
	"cutego/pkg/page"
	"cutego/refs"
	"github.com/go-xorm/xorm"
)

type LoginInfoDao struct {
}

// 查询公共sql
func (d LoginInfoDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_login_info")
}

// SelectPage 分页查询数据
func (d LoginInfoDao) SelectPage(query request.LoginInfoQuery) (*[]dataobject.SysLoginInfo, int64) {
	loginInfos := make([]dataobject.SysLoginInfo, 0)
	session := d.sql(refs.SqlDB.NewSession())
	session.And("user_name = ?", query.UserName)
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&loginInfos)
	if err != nil {
		common.ErrorLog(err)
		return nil, 0
	}
	return &loginInfos, total
}

// Insert 添加登录记录
func (d LoginInfoDao) Insert(body dataobject.SysLoginInfo) *dataobject.SysLoginInfo {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.Table("sys_login_info").Insert(&body)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
	return &body
}
