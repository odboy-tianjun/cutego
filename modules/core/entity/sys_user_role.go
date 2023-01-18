package entity

type SysUserRole struct {
	UserId int64 `xorm:"pk" json:"userId"` // 用户id
	RoleId int64 `xorm:"pk" json:"roleId"` // 角色id
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
