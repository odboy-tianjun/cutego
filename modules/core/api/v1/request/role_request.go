package request

import "cutego/pkg/base"

// RoleQuery 角色Get请求参数
type RoleQuery struct {
	base.GlobalQuery
	RoleName string `form:"roleName"` // 角色名称
	Status   string `form:"status"`   // 角色状态
	RoleKey  string `form:"roleKey"`  // 角色Key
}

// RoleBody 角色Post和Put参数
type RoleBody struct {
	RoleId int64  `json:"roleId"`
	Status string `json:"status"`
}
