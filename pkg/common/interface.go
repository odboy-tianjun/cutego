package common

// RedisData 存储数据结构
type RedisData struct {
	Key    string
	Field  string
	Value  string
	Expire int64
}

// RedisDataArray RedisData of array
type RedisDataArray []*RedisData

// IRedis redis client interface
type IRedis interface {
	// KEYS get patten key array
	KEYS(patten string) ([]string, error)

	// SCAN get patten key array
	SCAN(patten string) ([]string, error)

	// DEL delete k-v
	DEL(key string) (int, error)

	// DELALL delete key array
	DELALL(key []string) (int, error)

	// GET get k-v
	GET(key string) (string, error)

	// SET set k-v
	//SET(key string, value string) (int64, error)

	// SETEX set k-v expire seconds
	SETEX(key string, sec int, value string) (int64, error)

	// EXPIRE set key expire seconds
	EXPIRE(key string, sec int64) (int64, error)

	// HGETALL get map of key
	HGETALL(key string) (map[string]string, error)

	// HGET get value of key-field
	HGET(key string, field string) (string, error)

	// HSET set value of key-field
	//HSET(key string, field string, value string) (int64, error)

	// Write 向redis中写入多组数据
	Write(data RedisDataArray)
}
