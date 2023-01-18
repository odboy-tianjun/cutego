package dao

import (
	"cutego/modules/core/entity"
	"cutego/pkg/common"
	"cutego/refs"
)

type RoleMenuDao struct {
}

// Insert 添加角色菜单关系
func (d RoleMenuDao) Insert(list []entity.SysRoleMenu) int64 {
	session := refs.SqlDB.NewSession()
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
func (d RoleMenuDao) Delete(role entity.SysRole) {
	menu := entity.SysRoleMenu{
		RoleId: role.RoleId,
	}
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&menu)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
}
