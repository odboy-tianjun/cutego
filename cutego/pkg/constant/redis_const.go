package constant

const (
	redisErrorMsg  = "调用Redis发生异常, %s"
	redisDictKey   = "dict:"
	redisConfigKey = "config:"
)

// RedisConst Redis相关操作常量
type RedisConst struct{}

// GetRedisError Redis异常拼接常量
func (c RedisConst) GetRedisError() string {
	return redisErrorMsg
}

// GetRedisDictKey 获取Redis的Dict的key
func (c RedisConst) GetRedisDictKey() string {
	return redisDictKey
}

// GetRedisConfigKey 获取redis的config的key
func (c RedisConst) GetRedisConfigKey() string {
	return redisConfigKey
}
