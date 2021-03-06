package cache

import (
	"cutego/core/dao"
	"cutego/pkg/common"
	"cutego/pkg/constant"
)

// RemoveList 批量根据Key删除数据
// @Param list []string 键合集
func RemoveList(list []string) {
	dao.RedisDB.DELALL(list)
}

// RemoveKey 根据key删除
// @Param key 键
// @Return int 删除的数量
func RemoveCache(key string) int {
	del, err := dao.RedisDB.DEL(key)
	if err != nil {
		common.ErrorLog(err)
	}
	return del
}

// GetCache 获取缓存数据
// @Param key 键
// @Return string 值
func GetCache(key string) string {
	val, err := dao.RedisDB.GET(key)
	if err != nil {
		common.ErrorLog(constant.RedisConst{}.GetRedisError(), err.Error())
		return ""
	}
	return val
}

// SetCache 设置缓存数据
// @Param key 键
// @Param value 值
// @Return 新增的行数
func SetCache(key string, value interface{}) int {
	n, err := dao.RedisDB.SET(key, common.StructToJson(value))
	if err != nil {
		common.ErrorLog(constant.RedisConst{}.GetRedisError(), err.Error())
		return 0
	}
	return int(n)
}

// SetCache 设置缓存数据, 并指定过期时间
// @Param key 键
// @Param value 值
// @Param sec 过期时间(单位: 秒)
// @Return 新增的行数
func SetCacheTTL(key string, value interface{}, sec int) {
	dao.RedisDB.SETEX(key, sec, common.StructToJson(value))
}
