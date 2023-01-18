package cache

import (
	models2 "cutego/modules/core/entity"
	"cutego/pkg/cache"
	"cutego/pkg/common"
	"cutego/pkg/constant"
)

// GetRedisConfig 根据key从缓存中获取配置数据
// @Param key 键
// @Return *models2.SysConfig
func GetRedisConfig(key string) *models2.SysConfig {
	val := cache.GetCache(constant.RedisConst{}.GetRedisConfigKey() + key)
	s := new(models2.SysConfig)
	return common.JsonToStruct(val, s).(*models2.SysConfig)
}

// SetRedisConfig 将配置存入缓存
// @Param config models2.SysConfig
func SetRedisConfig(config models2.SysConfig) {
	cache.SetCache(config.ConfigKey, common.StructToJson(config))
}

// RemoveRedisConfig 从缓存中删除配置
// @Param configKey string 配置键
// @Return 影响的行数
func RemoveRedisConfig(configKey string) int {
	return cache.RemoveCache(constant.RedisConst{}.GetRedisConfigKey() + configKey)
}
