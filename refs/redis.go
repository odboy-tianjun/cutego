package refs

import redisTool "cutego/pkg/redispool"

// 配置redis数据库
func init() {
	RedisDB = redisTool.NewRedis()
}
