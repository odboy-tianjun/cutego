package service

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dao"
	"cutego/modules/core/dataobject"
)

type JobLogService struct {
	jobLogDao dao.JobLogDao
}

// FindPage 分页查询
func (s JobLogService) FindPage(query request.JobLogQuery) ([]dataobject.SysJobLog, int64) {
	return s.jobLogDao.SelectPage(query)
}

// Delete 批量删除
func (s JobLogService) Delete(ids []int64) bool {
	return s.jobLogDao.DeleteByIds(ids)
}

// Clean 清空所有日志
func (s JobLogService) Clean() bool {
	return s.jobLogDao.Clean()
}