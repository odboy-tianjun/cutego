package refs

import (
	"cutego/pkg/logger"
	redisTool "cutego/pkg/redispool"
)

// 配置redis数据库
func InitRedis() {
	logger.SugaredLogger.Infoln("redis init start...")
	RedisDB = redisTool.NewRedis()
	logger.SugaredLogger.Infoln("redis init end...")
}
