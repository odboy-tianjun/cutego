package service

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dao"
	"cutego/modules/core/dataobject"
)

type NoticeService struct {
	noticeDao dao.NoticeDao
}

// FindPage 查询集合数据
func (s NoticeService) FindPage(query request.NoticeQuery) (*[]dataobject.SysNotice, int64) {
	return s.noticeDao.SelectPage(query)
}

// Save 添加公告
func (s NoticeService) Save(notice dataobject.SysNotice) bool {
	return s.noticeDao.Insert(notice) > 0
}

// Remove 批量删除
func (s NoticeService) Remove(list []int64) bool {
	return s.noticeDao.Delete(list) > 0
}

// GetById 查询
func (s NoticeService) GetById(id int64) *dataobject.SysNotice {
	return s.noticeDao.SelectById(id)
}

// Edit 修改
func (s NoticeService) Edit(notice dataobject.SysNotice) bool {
	return s.noticeDao.Update(notice) > 0
}
