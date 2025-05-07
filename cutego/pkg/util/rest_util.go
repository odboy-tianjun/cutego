package util

import "cutego/pkg/page"

func NewPage(list []interface{}, total int64, pageSize int) page.Page {
	// 值对象, 不可变
	return page.Page{
		List:  list,
		Total: total,
		Size:  pageSize,
	}
}
