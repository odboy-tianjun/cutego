package util

import (
	"cutego/pkg/logging"
	"encoding/json"
)

// IF 三元表达式
func IF(condition bool, trueValue interface{}, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

// 结构体、Map等转Json字符串
// @Param v interface{}
// @Return Json字符串
func StructToJson(v interface{}) string {
	jsonBytes, err := json.Marshal(&v)
	if err != nil {
		logging.ErrorLog(err)
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
		logging.ErrorLog(err)
		return nil
	}
	//common.DebugLogf("JsonToStruct, obj=%v", s)
	return s
}
