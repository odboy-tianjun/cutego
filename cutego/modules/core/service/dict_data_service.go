package service

import (
	"cutego/modules/core/api/v1/request"
	cache2 "cutego/modules/core/cache"
	"cutego/modules/core/dao"
	"cutego/modules/core/dataobject"
	"cutego/pkg/constant"
)

type DictDataService struct {
	dictDataDao dao.DictDataDao
}

// FindByDictType 根据字典类型查询字典数据
func (s DictDataService) FindByDictType(dictType string) []dataobject.SysDictData {
	// 先从缓存中拉数据
	key := cache2.GetRedisDict(dictType)
	if key != nil {
		return key.([]dataobject.SysDictData)
	} else {
		// 缓存中为空, 从数据库中取数据
		return s.dictDataDao.SelectByDictType(dictType)
	}
}

// FindPage 查询字段数据集合
func (s DictDataService) FindPage(query request.DiceDataQuery) (*[]dataobject.SysDictData, int64) {
	return s.dictDataDao.SelectPage(query)
}

// GetByCode 根据code查询字典数据
// @Param code int64
// @Return *dataobject.SysDictData
func (s DictDataService) GetByCode(code int64) *dataobject.SysDictData {
	return s.dictDataDao.SelectByDictCode(code)
}

// Save 新增字典数据
// @Param data dataobject.SysDictData
// @Return bool
func (s DictDataService) Save(data dataobject.SysDictData) bool {
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
// @Return []dataobject.SysDictData
func (s DictDataService) GetNoCacheByType(dictType string) []dataobject.SysDictData {
	return s.dictDataDao.SelectByDictType(constant.RedisConst{}.GetRedisDictKey() + dictType)
}

// 修改字典数据
func (s DictDataService) Edit(data dataobject.SysDictData) bool {
	return s.dictDataDao.Update(data)
}
