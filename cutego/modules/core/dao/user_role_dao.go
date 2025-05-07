package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/pkg/logging"
	"cutego/refs"
)

type UserRoleDao struct {
}

// BatchInsert 批量新增用户角色信息
func (d UserRoleDao) BatchInsert(roles []dataobject.SysUserRole) {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.Table(dataobject.SysUserRole{}.TableName()).Insert(&roles)
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// Delete 删除用户和角色关系
func (d UserRoleDao) Delete(id int64) {
	role := dataobject.SysUserRole{
		UserId: id,
	}
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&role)
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// DeleteAuthUser 取消用户授权
func (d UserRoleDao) DeleteAuthUser(role dataobject.SysUserRole) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	i, err := session.Delete(&role)
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return i
}

// InsertAuthUsers 批量选择用户授权
func (d UserRoleDao) InsertAuthUsers(body request.UserRoleBody) int64 {
	ids := body.UserIds
	roles := make([]dataobject.SysUserRole, 0)
	for i := 0; i < len(ids); i++ {
		role := dataobject.SysUserRole{
			RoleId: body.RoleId,
			UserId: ids[i],
		}
		roles = append(roles, role)
	}
	session := refs.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&roles)
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}
