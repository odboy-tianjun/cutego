package excels

import (
	"reflect"
	"strings"
)

// 定义tagName名称
const tagName = "excel"

// Excels 定义Excel接口
type Excels interface {
	Excels(interface{}) (bool, string)
}

// DefaultExcels 默认Excel配置
type DefaultExcels struct {
}

// Excels 实现接口
func (e DefaultExcels) Excels(val interface{}) (bool, string) {
	return true, ""
}

// `excel:"name='',readConverterExp=''"`
// 获取标题头
func getExcelFromTagTitle(proName string, tag string) map[string]string {
	args := strings.SplitN(tag, ",", 2)
	m := make(map[string]string)
	for i := 0; i < len(args); i++ {
		contains := strings.Contains(args[i], "name")
		if contains {
			m[proName] = strings.Split(args[i], "=")[1]
			break
		}
	}
	return m
}

// `excel:"name='',readConverterExp='Y=1,N=2'"`
func getReadConverterExp(proName string, tag string) map[string][]map[string]string {
	list := make(map[string][]map[string]string, 0)
	args := strings.SplitN(tag, ",", 2)
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], "format") {
			maps := make([]map[string]string, 0)
			converter := strings.SplitN(args[i], "=", 2)[1]
			values := strings.Split(converter, ",")
			for j := 0; j < len(values); j++ {
				split := strings.Split(values[j], "=")
				m := make(map[string]string)
				m[split[0]] = split[1]
				maps = append(maps, m)
			}
			list[proName] = maps
			break
		}
	}
	return list
}

// ExcelCreate 获取name名称和readConverterExp集合
func ExcelCreate(s interface{}) ([]map[string]string, []map[string][]map[string]string) {
	titleList := make([]map[string]string, 0)
	expList := make([]map[string][]map[string]string, 0)
	// 获取值域
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	var title map[string]string
	var exp map[string][]map[string]string
	// 遍历反射树形
	for i := 0; i < t.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}
		name := t.Field(i).Name
		title = getExcelFromTagTitle(name, tag)
		exp = getReadConverterExp(name, tag)
		titleList = append(titleList, title)
		expList = append(expList, exp)
	}
	return titleList, expList
}
