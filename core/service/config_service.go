package service

import (
	"cutego/core/api/v1/request"
	cache2 "cutego/core/cache"
	"cutego/core/dao"
	"cutego/core/entity"
)

type ConfigService struct {
	configDao dao.ConfigDao
}

// GetConfigKey 根据键名查询参数配置信息
func (s ConfigService) GetConfigKey(param string) *entity.SysConfig {
	// 从缓存中取出数据判断是否存在, 存在直接使用, 不存在就从数据库查询
	val := cache2.GetRedisConfig(param)
	if val != nil {
		return val
	}
	configKey := s.configDao.SelectByConfigKey(param)
	cache2.SetRedisConfig(*configKey)
	return configKey
}

// FindPage 分页查询数据
func (s ConfigService) FindPage(query request.ConfigQuery) (*[]entity.SysConfig, int64) {
	return s.configDao.SelectPage(query)
}

// CheckConfigKeyUnique 校验是否存在
func (s ConfigService) CheckConfigKeyUnique(config entity.SysConfig) bool {
	return s.configDao.CheckConfigKeyUnique(config) > 0
}

// Save 添加数据
func (s ConfigService) Save(config entity.SysConfig) int64 {
	return s.configDao.Insert(config)
}

// GetInfo 查询数据
func (s ConfigService) GetInfo(id int64) *entity.SysConfig {
	return s.configDao.SelectById(id)
}

// Edit 修改数据
func (s ConfigService) Edit(config entity.SysConfig) int64 {
	return s.configDao.Update(config)
}

// Remove 批量删除
func (s ConfigService) Remove(list []int64) bool {
	return s.configDao.Delete(list)
}

// CheckConfigByIds 根据id集合查询
func (s ConfigService) CheckConfigByIds(list []int64) *[]entity.SysConfig {
	return s.configDao.CheckConfigByIds(list)
}

// FindAll 查询所有数据
func (s ConfigService) FindAll() *[]entity.SysConfig {
	return s.configDao.SelectAll()
}
