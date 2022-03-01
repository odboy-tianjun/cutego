package dao

import (
	models2 "cutego/core/entity"
	"cutego/pkg/common"
)

type RoleMenuDao struct {
}

// Insert 添加角色菜单关系
func (d RoleMenuDao) Insert(list []models2.SysRoleMenu) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&list)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
	return insert
}

// Delete 删除角色和菜单关系
func (d RoleMenuDao) Delete(role models2.SysRole) {
	menu := models2.SysRoleMenu{
		RoleId: role.RoleId,
	}
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&menu)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
}
