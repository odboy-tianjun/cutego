package cache

import (
	"cutego/pkg/constant"
	"cutego/refs"
	"encoding/json"
	"fmt"
)

// RemoveList 批量根据Key删除数据
// @Param list []string 键合集
func RemoveList(list []string) {
	refs.RedisDB.DELALL(list)
}

// RemoveKey 根据key删除
// @Param key 键
// @Return int 删除的数量
func RemoveCache(key string) int {
	del, err := refs.RedisDB.DEL(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	return del
}

// GetCache 获取缓存数据
// @Param key 键
// @Return string 值
func GetCache(key string) string {
	val, err := refs.RedisDB.GET(key)
	if err != nil {
		fmt.Println(constant.RedisConst{}.GetRedisError(), err.Error())
		return ""
	}
	return val
}

// SetCache 设置缓存数据
// @Param key 键
// @Param value 值
// @Return 新增的行数
func SetCache(key string, value interface{}) int {
	n, err := refs.RedisDB.SET(key, StructToJson(value))
	if err != nil {
		fmt.Println(constant.RedisConst{}.GetRedisError(), err.Error())
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
	refs.RedisDB.SETEX(key, sec, StructToJson(value))
}

// 结构体、Map等转Json字符串
// @Param v interface{}
// @Return Json字符串
func StructToJson(v interface{}) string {
	jsonBytes, err := json.Marshal(&v)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	s := string(jsonBytes)
	//DebugLogf("StructToJson, json=%s", s)
	return s
}

// Json字符串转结构体、Map等
//
// 单个对象
// s := new(models2.SysConfig)
// return common.JsonToStruct(get, s).(*models2.SysConfig)
//
// 切片(interface{}.(期望类型))
// s := make([]interface {}, 0)
// target := common.JsonToStruct(get, s)
// target.([]dataobject.SysDictData)
//
// @Param data Json字符串
// @Param s 容器(结构体、Map等)
// @Return interface{}
func JsonToStruct(data string, s interface{}) interface{} {
	err := json.Unmarshal([]byte(data), &s)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	//common.DebugLogf("JsonToStruct, obj=%v", s)
	return s
}
