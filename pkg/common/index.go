package common

import (
	"cutego/pkg/config"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// IntToString int转string
func IntToString(n int) string {
	return strconv.Itoa(n)
}

// StringToInt string转int
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// StringToInt64 string转int64
func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

// Int64ToString int64转string
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func mapToBytes(data map[string]interface{}) []byte {
	bytes, _ := json.Marshal(data)
	return bytes
}

// MapToStruct map转struct
func MapToStruct(data map[string]interface{}, v interface{}) {
	_ = json.Unmarshal(mapToBytes(data), v)
}

// GetDirPath 获取目录路径
func GetDirPath(resType string) string {
	sysType := runtime.GOOS
	switch sysType {
	case "linux":
		if resType == "log" {
			return config.AppCoreConfig.CuteGoConfig.File.Linux.Logs
		} else if resType == "avatar" {
			return config.AppCoreConfig.CuteGoConfig.File.Linux.Avatar
		} else if resType == "file" {
			return config.AppCoreConfig.CuteGoConfig.File.Linux.Path
		}
		break
	case "windows":
		if resType == "log" {
			return config.AppCoreConfig.CuteGoConfig.File.Windows.Logs
		} else if resType == "avatar" {
			return config.AppCoreConfig.CuteGoConfig.File.Windows.Avatar
		} else if resType == "file" {
			return config.AppCoreConfig.CuteGoConfig.File.Windows.Path
		}
		break
	case "mac":
		if resType == "log" {
			return config.AppCoreConfig.CuteGoConfig.File.Mac.Logs
		} else if resType == "avatar" {
			return config.AppCoreConfig.CuteGoConfig.File.Mac.Avatar
		} else if resType == "file" {
			return config.AppCoreConfig.CuteGoConfig.File.Mac.Path
		}
		break
	case "darwin":
		if resType == "log" {
			return config.AppCoreConfig.CuteGoConfig.File.Mac.Logs
		} else if resType == "avatar" {
			return config.AppCoreConfig.CuteGoConfig.File.Mac.Avatar
		} else if resType == "file" {
			return config.AppCoreConfig.CuteGoConfig.File.Mac.Path
		}
	}
	return config.AppCoreConfig.CuteGoConfig.File.Linux.Logs
}

// CreateAllDir 递归创建文件夹
func CreateAllDir(filePath string) error {
	if !IsFileOrDirExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败, error info:", err)
			return err
		}
		return err
	}
	return nil
}

// IsFileOrDirExist 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsFileOrDirExist(path string) bool {
	// os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 类三元表达式
// condition 成立条件
// trueVal 当条件为true时返回
// false 当条件为false时返回
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// 结构体、Map等转Json字符串
// @Param v interface{}
// @Return Json字符串
func StructToJson(v interface{}) string {
	jsonBytes, err := json.Marshal(&v)
	if err != nil {
		ErrorLog(err)
		return ""
	}
	s := string(jsonBytes)
	DebugLogf("StructToJson, json=%s", s)
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
		ErrorLog(err)
		return nil
	}
	DebugLogf("JsonToStruct, obj=%v", s)
	return s
}
