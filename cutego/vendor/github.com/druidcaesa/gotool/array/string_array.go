package array

import (
	"github.com/druidcaesa/gotool/logs"
	"strconv"
)

type StrArray struct {
	logs logs.Logs
}

// StringToInt64 String array to int64 array
func (sa *StrArray) StringToInt64(s []string) ([]int64, error) {
	int64s := make([]int64, len(s))
	for i, item := range s {
		parseInt, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			sa.logs.ErrorLog().Println(err.Error())
			return nil, err
		}
		int64s[i] = parseInt
	}
	return int64s, nil
}

// StringToInt32 String array to int32 array
func (sa *StrArray) StringToInt32(s []string) ([]int32, error) {
	int32s := make([]int32, len(s))
	for i, item := range s {
		parseInt, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			sa.logs.ErrorLog().Println(err.Error())
			return nil, err
		}
		int32s[i] = int32(parseInt)
	}
	return int32s, nil
}

// ArrayDuplication Array deduplication String数组去重
func (sa *StrArray) ArrayDuplication(arr []string) []string {
	var out []string
	tmp := make(map[string]byte)
	for _, v := range arr {
		tmplen := len(tmp)
		tmp[v] = 0
		if len(tmp) != tmplen {
			out = append(out, v)
		}
	}
	return out
}
