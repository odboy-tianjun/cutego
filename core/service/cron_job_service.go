package service

import (
	"cutego/core/api/v1/request"
	"cutego/core/dao"
	"cutego/core/entity"
)

type CronJobService struct {
	cronJobService dao.CronJobDao
}

// FindPage 分页查询数据
func (s CronJobService) FindPage(query request.CronJobQuery) ([]entity.SysCronJob, int64) {
	return s.cronJobService.SelectPage(query)
}

// Save 添加数据
func (s CronJobService) Save(config entity.SysCronJob) int64 {
	return s.cronJobService.Insert(config)
}

// GetInfo 查询数据
func (s CronJobService) GetInfo(id int64) *entity.SysCronJob {
	return s.cronJobService.SelectById(id)
}

// GetInfo 查询数据
func (s CronJobService) GetInfoByAlias(funcAlias string) *entity.SysCronJob {
	return s.cronJobService.SelectByFuncAlias(funcAlias)
}

// Edit 修改数据
func (s CronJobService) Edit(config entity.SysCronJob) int64 {
	return s.cronJobService.Update(config)
}

// Remove 批量删除
func (s CronJobService) Remove(list []int64) bool {
	return s.cronJobService.Delete(list)
}
