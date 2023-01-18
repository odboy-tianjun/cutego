package service

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dao"
	"cutego/modules/core/dataobject"
)

type CronJobService struct {
	cronJobDao dao.CronJobDao
}

// FindPage 分页查询数据
func (s CronJobService) FindPage(query request.CronJobQuery) ([]dataobject.SysCronJob, int64) {
	return s.cronJobDao.SelectPage(query)
}

// Save 添加数据
func (s CronJobService) Save(config dataobject.SysCronJob) int64 {
	return s.cronJobDao.Insert(config)
}

// GetInfo 查询数据
func (s CronJobService) GetInfo(id int64) *dataobject.SysCronJob {
	return s.cronJobDao.SelectById(id)
}

// GetInfoByAlias 查询数据
func (s CronJobService) GetInfoByAlias(funcAlias string) *dataobject.SysCronJob {
	return s.cronJobDao.SelectByFuncAlias(funcAlias)
}

// Edit 修改数据
func (s CronJobService) Edit(config dataobject.SysCronJob) int64 {
	return s.cronJobDao.Update(config)
}

// Remove 批量删除
func (s CronJobService) Remove(list []int64) bool {
	return s.cronJobDao.Delete(list)
}

// FindAll 查找所有
func (s CronJobService) FindAll() ([]dataobject.SysCronJob, int) {
	return s.cronJobDao.SelectAll()
}
