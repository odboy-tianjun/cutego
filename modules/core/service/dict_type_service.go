package service

import (
	"cutego/modules/core/api/v1/request"
	cache2 "cutego/modules/core/cache"
	"cutego/modules/core/dao"
	"cutego/modules/core/entity"
)

type DictTypeService struct {
	dictTypeDao dao.DictTypeDao
	dictDataDao dao.DictDataDao
}

// FindPage 分页查询字典类型数据
func (s DictTypeService) FindPage(query request.DictTypeQuery) (*[]entity.SysDictType, int64) {
	return s.dictTypeDao.SelectPage(query)
}

// GetById 根据id查询字典类型数据
func (s DictTypeService) GetById(id int64) *entity.SysDictType {
	return s.dictTypeDao.SelectById(id)
}

// CheckDictTypeUnique 检验字典类型是否存在
func (s DictTypeService) CheckDictTypeUnique(dictType entity.SysDictType) bool {
	return s.dictTypeDao.CheckDictTypeUnique(dictType) > 0
}

// Edit 修改字典数据
func (s DictTypeService) Edit(dictType entity.SysDictType) bool {
	return s.dictTypeDao.Update(dictType)
}

// Save 新增字典类型
func (s DictTypeService) Save(dictType entity.SysDictType) bool {
	insert := s.dictTypeDao.Insert(dictType)
	if insert > 0 {
		cache2.SetRedisDict(dictType.DictType, nil)
	}
	return insert > 0
}

// Remove 批量删除
func (s DictTypeService) Remove(ids []int64) bool {
	return s.dictTypeDao.Delete(ids)
}

// FindAll 查询所有字典类型数据
func (s DictTypeService) FindAll() []*entity.SysDictType {
	return s.dictTypeDao.SelectAll()
}

// RemoveAllCache 删除所有字典缓存
func (s DictTypeService) RemoveAllCache() []string {
	typeList := make([]string, 0)
	allType := s.FindAll()
	for _, dictType := range allType {
		typeList = append(typeList, dictType.DictType)
	}
	// 删除缓存
	cache2.RemoveRedisDictList(typeList)
	return typeList
}

// LoadDictCache 将字典数据存入缓存
func (s DictTypeService) LoadDictCache() {
	typeList := make([]string, 0)
	allType := s.FindAll()
	for _, dictType := range allType {
		typeList = append(typeList, dictType.DictType)
	}
	allData := s.dictDataDao.GetDiceDataAll()
	for _, key := range typeList {
		list := make([]entity.SysDictData, 0)
		for _, data := range *allData {
			if key == data.DictType {
				list = append(list, data)
			}
		}
		cache2.SetRedisDict(key, list)
	}
}

// RefreshCache 刷新缓存数据
func (s DictTypeService) RefreshCache() {
	typeList := s.RemoveAllCache()
	allData := s.dictDataDao.GetDiceDataAll()
	for _, key := range typeList {
		list := make([]entity.SysDictData, 0)
		for _, data := range *allData {
			if key == data.DictType {
				list = append(list, data)
			}
		}
		cache2.SetRedisDict(key, list)
	}
}
