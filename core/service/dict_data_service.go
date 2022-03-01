package service

import (
	"cutego/core/api/v1/request"
	cache2 "cutego/core/cache"
	"cutego/core/dao"
	"cutego/core/entity"
	"cutego/pkg/constant"
)

type DictDataService struct {
	dictDataDao dao.DictDataDao
}

// FindByDictType 根据字典类型查询字典数据
func (s DictDataService) FindByDictType(dictType string) []entity.SysDictData {
	// 先从缓存中拉数据
	key := cache2.GetRedisDict(dictType)
	if key != nil {
		return key.([]entity.SysDictData)
	} else {
		// 缓存中为空, 从数据库中取数据
		return s.dictDataDao.SelectByDictType(dictType)
	}
}

// FindPage 查询字段数据集合
func (s DictDataService) FindPage(query request.DiceDataQuery) (*[]entity.SysDictData, int64) {
	return s.dictDataDao.SelectPage(query)
}

// GetByCode 根据code查询字典数据
// @Param code int64
// @Return *entity.SysDictData
func (s DictDataService) GetByCode(code int64) *entity.SysDictData {
	return s.dictDataDao.SelectByDictCode(code)
}

// Save 新增字典数据
// @Param data entity.SysDictData
// @Return bool
func (s DictDataService) Save(data entity.SysDictData) bool {
	insert := s.dictDataDao.Insert(data)
	if insert > 0 {
		// 刷新缓存数据
		byType := s.GetNoCacheByType(data.DictType)
		cache2.SetRedisDict(data.DictType, byType)
	}
	return insert > 0
}

// Remove 删除数据
// @Param codes 字典code集合
// @Return bool
func (s DictDataService) Remove(codes []int64) bool {
	dictType := s.GetByCode(codes[0]).DictType
	remove := s.dictDataDao.Delete(codes)
	if remove {
		// 刷新缓存
		code := s.GetNoCacheByType(dictType)
		cache2.SetRedisDict(dictType, code)
	}
	return remove
}

// GetNoCacheByType 根据字典类型查询字典数据
// @Param dictType 字典类型
// @Return []entity.SysDictData
func (s DictDataService) GetNoCacheByType(dictType string) []entity.SysDictData {
	return s.dictDataDao.SelectByDictType(constant.RedisConst{}.GetRedisDictKey() + dictType)
}

// 修改字典数据
func (s DictDataService) Edit(data entity.SysDictData) bool {
	return s.dictDataDao.Update(data)
}
