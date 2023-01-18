package entity

// SysRoleMenu 角色菜单数据结构
type SysRoleMenu struct {
	RoleId int64 `xorm:"pk" json:"roleId"` // 角色id
	MenuId int64 `xorm:"pk" json:"menuId"` // 菜单id
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
