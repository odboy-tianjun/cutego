package cache

import (
	"cutego/pkg/cache"
	"cutego/pkg/config"
	"cutego/pkg/constant"
)

// SetRedisToken 将token存入到redis
// @Param username string 用户名
// @Param token string token令牌
func SetRedisToken(username string, token string) {
	if config.AppEnvConfig.Login.Single {
		cache.SetCacheTTL(constant.RedisOnlineUserKey+username, token, 3600)
	}
}

// RemoveRedisToken 将token从redis中删除
// @Param username string 用户名
// @Return 删除的行数
func RemoveRedisToken(username string) int {
	// 不管是不是单点登录, 直接踢
	return cache.RemoveCache(constant.RedisOnlineUserKey + username)
}
