package request

// UserRoleBody 用户角色Post和Put参数接收结构体
type UserRoleBody struct {
	RoleId  int64   `form:"roleId" json:"roleId"`
	UserIds []int64 `form:"userIds" json:"userIds"`
}
