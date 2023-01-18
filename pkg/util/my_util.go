package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// IF 三元表达式
func IF(condition bool, trueValue interface{}, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

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
