package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/pkg/logging"
	"cutego/pkg/page"
	"cutego/refs"
	"github.com/druidcaesa/gotool"
)

type NoticeDao struct {
}

// SelectPage 查询集合
func (d NoticeDao) SelectPage(query request.NoticeQuery) (*[]dataobject.SysNotice, int64) {
	notices := make([]dataobject.SysNotice, 0)
	session := refs.SqlDB.NewSession().Table(dataobject.SysNotice{}.TableName())
	if gotool.StrUtils.HasNotEmpty(query.NoticeTitle) {
		session.And("notice_title like concat('%', ?, '%')", query.NoticeTitle)
	}
	if gotool.StrUtils.HasNotEmpty(query.NoticeType) {
		session.And("notice_type = ?", query.NoticeType)
	}
	if gotool.StrUtils.HasNotEmpty(query.CreateBy) {
		session.And("create_by like concat('%', ?, '%')", query.CreateBy)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&notices)
	if err != nil {
		logging.ErrorLog(err)
		return nil, 0
	}
	return &notices, total
}

// Insert 添加数据
func (d NoticeDao) Insert(notice dataobject.SysNotice) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&notice)
	if err != nil {
		session.Rollback()
		logging.ErrorLog(err)
		return 0
	}
	session.Commit()
	return insert
}

// Delete 批量删除
func (d NoticeDao) Delete(list []int64) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	i, err := session.In("notice_id", list).Delete(&dataobject.SysNotice{})
	if err != nil {
		session.Rollback()
		logging.ErrorLog(err)
		return 0
	}
	session.Commit()
	return i
}

// SelectById 查询
func (d NoticeDao) SelectById(id int64) *dataobject.SysNotice {
	notice := dataobject.SysNotice{}
	_, err := refs.SqlDB.NewSession().Where("notice_id = ?", id).Get(&notice)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &notice
}

// Update 修改数据
func (d NoticeDao) Update(notice dataobject.SysNotice) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("notice_id = ?", notice.NoticeId).Update(&notice)
	if err != nil {
		session.Rollback()
		logging.ErrorLog(err)
		return 0
	}
	session.Commit()
	return update
}
