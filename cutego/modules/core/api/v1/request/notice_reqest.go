package request

import "cutego/pkg/base"

type NoticeQuery struct {
	base.GlobalQuery
	NoticeTitle string `form:"noticeTitle"`
	NoticeType  string `form:"noticeType"`
	CreateBy    string `form:"createBy"`
}
