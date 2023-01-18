package dao

import (
	"cutego/modules/core/dataobject"
	"cutego/pkg/logging"
	"cutego/refs"
)

type RoleMenuDao struct {
}

// Insert 添加角色菜单关系
func (d RoleMenuDao) Insert(list []dataobject.SysRoleMenu) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&list)
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
	return insert
}

// Delete 删除角色和菜单关系
func (d RoleMenuDao) Delete(role dataobject.SysRole) {
	menu := dataobject.SysRoleMenu{
		RoleId: role.RoleId,
	}
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&menu)
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
}
