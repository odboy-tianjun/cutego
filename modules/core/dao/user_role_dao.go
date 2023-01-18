package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/entity"
	"cutego/pkg/common"
)

type UserRoleDao struct {
}

// BatchInsert 批量新增用户角色信息
func (d UserRoleDao) BatchInsert(roles []entity.SysUserRole) {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Table(entity.SysUserRole{}.TableName()).Insert(&roles)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// Delete 删除用户和角色关系
func (d UserRoleDao) Delete(id int64) {
	role := entity.SysUserRole{
		UserId: id,
	}
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&role)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// DeleteAuthUser 取消用户授权
func (d UserRoleDao) DeleteAuthUser(role entity.SysUserRole) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	i, err := session.Delete(&role)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return i
}

// InsertAuthUsers 批量选择用户授权
func (d UserRoleDao) InsertAuthUsers(body request.UserRoleBody) int64 {
	ids := body.UserIds
	roles := make([]entity.SysUserRole, 0)
	for i := 0; i < len(ids); i++ {
		role := entity.SysUserRole{
			RoleId: body.RoleId,
			UserId: ids[i],
		}
		roles = append(roles, role)
	}
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&roles)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}
