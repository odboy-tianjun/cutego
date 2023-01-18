package service

import (
	"cutego/core/api/v1/request"
	"cutego/core/dao"
	"cutego/core/entity"
)

type CronJobService struct {
	cronJobDao dao.CronJobDao
}

// FindPage 分页查询数据
func (s CronJobService) FindPage(query request.CronJobQuery) ([]entity.SysCronJob, int64) {
	return s.cronJobDao.SelectPage(query)
}

// Save 添加数据
func (s CronJobService) Save(config entity.SysCronJob) int64 {
	return s.cronJobDao.Insert(config)
}

// GetInfo 查询数据
func (s CronJobService) GetInfo(id int64) *entity.SysCronJob {
	return s.cronJobDao.SelectById(id)
}

// GetInfoByAlias 查询数据
func (s CronJobService) GetInfoByAlias(funcAlias string) *entity.SysCronJob {
	return s.cronJobDao.SelectByFuncAlias(funcAlias)
}

// Edit 修改数据
func (s CronJobService) Edit(config entity.SysCronJob) int64 {
	return s.cronJobDao.Update(config)
}

// Remove 批量删除
func (s CronJobService) Remove(list []int64) bool {
	return s.cronJobDao.Delete(list)
}

// FindAll 查找所有
func (s CronJobService) FindAll() ([]entity.SysCronJob, int) {
	return s.cronJobDao.SelectAll()
}
