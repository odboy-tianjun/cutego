package util

import (
	"bytes"
	"cutego/pkg/common"
	"encoding/gob"
)

// DeepCopy 深度拷贝对象
// @Param src 原对象
// @Param dst 目标对象
// @Return 目标对象
// @Usage util.DeepCopy(item, &response.CronJobPageResponse{})
//
// @Author tianjun@odboy.cn
// @Date 2022-03-01
func DeepCopy(src, dst interface{}) interface{} {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	err := gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
	if err != nil {
		common.ErrorLogf("src(%v)---> dst(%v), error", src, dst)
	}
	return dst
}
