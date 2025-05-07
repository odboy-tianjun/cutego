package request

import "cutego/pkg/base"

type LogQuery struct {
	base.GlobalQuery
	Uid     string `form:"uid"`
	Content string `form:"content"`
}
