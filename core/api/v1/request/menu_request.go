package request

import "cutego/pkg/base"

// MenuQuery 菜单查询条件封装
type MenuQuery struct {
	base.GlobalQuery
	MenuName string `json:"menuName" form:"menuName"`
	Visible  string `json:"visible" form:"visible"`
	Status   string `json:"status" form:"status"`
	UserId   int64  `json:"userId" form:"userId"`
}
