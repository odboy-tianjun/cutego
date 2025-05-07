package redispool

import (
	"cutego/pkg/common"
	"cutego/pkg/config"
	"cutego/pkg/logging"
	"fmt"
	"github.com/druidcaesa/gotool"
	"github.com/gomodule/redigo/redis"
	"time"
)

// https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples
// https://github.com/gomodule/redigo

// RedisClient redis client instance
type RedisClient struct {
	pool *redis.Pool
	// 数据接收
	chanRx chan common.RedisDataArray
	// 是否退出
	isExit bool
}

// NewRedis new redis client
func NewRedis() *RedisClient {
	return &RedisClient{
		pool:   newPool(),
		chanRx: make(chan common.RedisDataArray, 100),
	}
}

// newPool 线程池
func newPool() *redis.Pool {
	if config.AppEnvConfig.Redis.Pool.MaxIdle == 0 {
		config.AppEnvConfig.Redis.Pool.MaxIdle = 3
	}
	return &redis.Pool{
		MaxIdle:     config.AppEnvConfig.Redis.Pool.MaxIdle,
		IdleTimeout: time.Duration(config.AppEnvConfig.Redis.Pool.MaxWait) * time.Second,
		MaxActive:   config.AppEnvConfig.Redis.Pool.MaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.AppEnvConfig.Redis.Host, config.AppEnvConfig.Redis.Port))
			if err != nil {
				logging.FatalfLog("Redis.Dial: %v", err)
				return nil, err
			}
			if gotool.StrUtils.HasNotEmpty(config.AppEnvConfig.Redis.Password) {
				if _, err := c.Do("AUTH", config.AppEnvConfig.Redis.Password); err != nil {
					c.Close()
					logging.FatalfLog("Redis.AUTH: %v", err)
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", config.AppEnvConfig.Redis.Database); err != nil {
				c.Close()
				logging.FatalfLog("Redis.SELECT: %v", err)
				return nil, err
			}
			return c, nil
		},
	}
}

// Start 启动接收任务协程
func (r *RedisClient) Start() {
	r.isExit = false
	// 开启协程用于循环接收数据
	go r.loopRead()
}

// Stop 停止接收任务
func (r *RedisClient) Stop() {
	r.isExit = true
	// 关闭数据接收通道
	close(r.chanRx)
	// 关闭redis线程池
	r.pool.Close()
}

// Write 向redis中写入多组数据
func (r *RedisClient) Write(data common.RedisDataArray) {
	r.chanRx <- data
}

// loopRead 循环接收数据
func (r *RedisClient) loopRead() {
	for !r.isExit {
		select {
		case rx := <-r.chanRx:
			for _, it := range rx {
				if len(it.Key) > 0 {
					if len(it.Field) > 0 {
						if _, err := r.HSET(it.Key, it.Field, it.Value); err != nil {
							logging.DebugLogf("[%s, %s, %s]: %s\n", it.Key, it.Field, it.Value, err.Error())
						}
					} else {
						if _, err := r.SET(it.Key, it.Value); err != nil {
							logging.DebugLogf("[%s, %s, %s]: %s\n", it.Key, it.Field, it.Value, err.Error())
						}
					}
					if it.Expire > 0 {
						r.EXPIRE(it.Key, it.Expire)
					}
				}
			}
		}
	}

}

// Error get redis connect error
func (r *RedisClient) Error() error {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Err()
}

// 常用Redis操作命令的封装
// http://redis.io/commands

// KEYS get patten key array
func (r *RedisClient) KEYS(patten string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", patten))
}

// SCAN 获取大量key
func (r *RedisClient) SCAN(patten string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	var out []string
	var cursor uint64 = 0xffffff
	isFirst := true
	for cursor != 0 {
		if isFirst {
			cursor = 0
			isFirst = false
		}
		arr, err := conn.Do("SCAN", cursor, "MATCH", patten, "COUNT", 100)
		if err != nil {
			return out, err
		}
		switch arr := arr.(type) {
		case []interface{}:
			cursor, _ = redis.Uint64(arr[0], nil)
			it, _ := redis.Strings(arr[1], nil)
			out = append(out, it...)
		}
	}
	out = gotool.StrArrayUtils.ArrayDuplication(out)
	return out, nil
}

// DEL delete k-v
func (r *RedisClient) DEL(key string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("DEL", key))
}

// DELALL delete key array
func (r *RedisClient) DELALL(key []string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	arr := make([]interface{}, len(key))
	for i, v := range key {
		arr[i] = v
	}
	return redis.Int(conn.Do("DEL", arr...))
}

// GET get k-v
func (r *RedisClient) GET(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

// SET set k-v
func (r *RedisClient) SET(key string, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("SET", key, value))
}

// SETEX set k-v expire seconds
func (r *RedisClient) SETEX(key string, sec int, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("SETEX", key, sec, value))
}

// EXPIRE set key expire seconds
func (r *RedisClient) EXPIRE(key string, sec int64) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("EXPIRE", key, sec))
}

// HGETALL get map of key
func (r *RedisClient) HGETALL(key string) (map[string]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", key))
}

// HGET get value of key-field
func (r *RedisClient) HGET(key string, field string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("HGET", key, field))
}

// HSET set value of key-field
func (r *RedisClient) HSET(key string, field string, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("HSET", key, field, value))
}
