package refs

import (
	"cutego/pkg/logging"
	redisTool "cutego/pkg/redispool"
)

// 配置redis数据库
func init() {
	logging.InfoLog("redis init start...")
	RedisDB = redisTool.NewRedis()
	logging.InfoLog("redis init end...")
}
