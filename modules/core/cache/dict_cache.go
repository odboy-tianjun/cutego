package cache

import (
	models2 "cutego/modules/core/dataobject"
	"cutego/pkg/cache"
	"cutego/pkg/common"
	"cutego/pkg/constant"
)

// GetRedisDict 根据key获取缓存中的字典数据
// @Param key string 键
// @Return interface {}
func GetRedisDict(key string) interface{} {
	val := cache.GetCache(key)
	s := make([]interface{}, 0)
	return common.JsonToStruct(val, s)
}

// SetRedisDict 保存字典数据
// @Param dictType string 字典类型
// @Param list []models2.SysDictData
func SetRedisDict(dictType string, list []models2.SysDictData) {
	cache.SetCache(constant.RedisConst{}.GetRedisDictKey()+dictType, list)
}

// RemoveRedisDictList 批量删除字典数据
// @Param dictType []string 字典类型集合
func RemoveRedisDictList(dictType []string) {
	includeKey := make([]string, 0)
	header := constant.RedisConst{}.GetRedisDictKey()
	for _, e := range dictType {
		includeKey = append(includeKey, header+e)
	}
	cache.RemoveList(includeKey)
}
