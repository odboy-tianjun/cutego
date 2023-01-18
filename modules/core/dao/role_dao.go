package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/entity"
	"cutego/pkg/common"
	"cutego/pkg/page"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
)

type RoleDao struct {
}

// 角色公用sql
func (d RoleDao) sqlSelectJoin() *xorm.Session {
	return SqlDB.Table([]string{entity.SysRole{}.TableName(), "r"}).
		Join("LEFT", []string{"sys_user_role", "ur"}, "ur.role_id = r.role_id").
		Join("LEFT", []string{"sys_user", "u"}, "u.user_id = ur.user_id").
		Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
}

// 用户角色关系查询sql
func (d RoleDao) sqlSelectRoleAndUser() *xorm.Session {
	return SqlDB.Table([]string{entity.SysRole{}.TableName(), "r"}).
		Join("LEFT", []string{"sys_user_role", "ur"}, "ur.role_id = r.role_id").
		Join("LEFT", []string{"sys_user", "u"}, "u.user_id = ur.user_id")
}

// SelectPage 根据条件查询角色数据
func (d RoleDao) SelectPage(q *request.RoleQuery) ([]*entity.SysRole, int64) {
	roles := make([]*entity.SysRole, 0)
	session := d.sqlSelectJoin()
	if !gotool.StrUtils.HasEmpty(q.RoleName) {
		session.And("r.role_name like concat('%', ?, '%')", q.RoleName)
	}
	if !gotool.StrUtils.HasEmpty(q.Status) {
		session.And("r.status = ?", q.Status)
	}
	if !gotool.StrUtils.HasEmpty(q.RoleKey) {
		session.And("r.role_key like concat('%', ?, '%')", q.RoleKey)
	}
	if !gotool.StrUtils.HasEmpty(q.BeginTime) {
		session.And("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d')", q.BeginTime)
	}
	if !gotool.StrUtils.HasEmpty(q.EndTime) {
		session.And("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d')", q.EndTime)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(q.PageSize, page.StartSize(q.PageNum, q.PageSize)).OrderBy("r.role_sort").Find(&roles)
	if err != nil {
		return nil, 0
	}
	return roles, total
}

// SelectAll 查询所有角色
func (d RoleDao) SelectAll() []*entity.SysRole {
	sql := d.sqlSelectJoin()
	roles := make([]*entity.SysRole, 0)
	err := sql.Find(&roles)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return roles
}

// SelectRoleListByUserId 根据用户id查询用户角色id集合
func (d RoleDao) SelectRoleListByUserId(userId int64) *[]int64 {
	sqlSelectRoleAndUser := d.sqlSelectRoleAndUser()
	var roleIds []int64
	err := sqlSelectRoleAndUser.Cols("r.role_id").Where("u.user_id = ?", userId).Find(&roleIds)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &roleIds
}

// SelectRolePermissionByUserId 查询用户角色集合
func (d RoleDao) SelectRolePermissionByUserId(id int64) *[]string {
	var roleKeys []string
	err := d.sqlSelectJoin().Cols("r.role_key").Where("r.del_flag = '0'").And("ur.user_id = ?", id).Find(&roleKeys)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &roleKeys
}

// GetRoleListByUserId 根据用户ID查询角色
func (d RoleDao) GetRoleListByUserId(id int64) *[]entity.SysRole {
	roles := make([]entity.SysRole, 0)
	err := d.sqlSelectJoin().Where("r.del_flag = '0'").And("ur.user_id = ?", id).Find(&roles)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &roles
}

// SelectRoleByRoleId 根据角色id查询角色数据
func (d RoleDao) SelectRoleByRoleId(id int64) *entity.SysRole {
	role := entity.SysRole{}
	_, err := d.sqlSelectJoin().Where("r.role_id = ?", id).Get(&role)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &role
}

// CheckRoleNameUnique 校验角色名称是否唯一
func (d RoleDao) CheckRoleNameUnique(role entity.SysRole) int64 {
	session := SqlDB.Table(role.TableName()).Where("role_name = ?", role.RoleName)
	if role.RoleId > 0 {
		session.And("role_id != ?", role.RoleId)
	}
	count, err := session.Count(&role)
	if err != nil {
		common.ErrorLog(err)
	}
	return count
}

// CheckRoleKeyUnique 校验角色权限是否唯一
func (d RoleDao) CheckRoleKeyUnique(role entity.SysRole) int64 {
	session := SqlDB.Table(role.TableName()).Where("role_key = ?", role.RoleKey)
	if role.RoleId > 0 {
		session.And("role_id != ?", role.RoleId)
	}
	count, err := session.Count(&role)
	if err != nil {
		common.ErrorLog(err)
	}
	return count
}

// Add 添加角色进入数据库操作
func (d RoleDao) Insert(role entity.SysRole) entity.SysRole {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Insert(&role)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
	return role
}

// Update 修改数据
func (d RoleDao) Update(role entity.SysRole) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("role_id = ?", role.RoleId).Update(&role)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return update
}

// Delete 删除角色
func (d RoleDao) Delete(role entity.SysRole) int64 {
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

// UpdateRoleStatus 修改角色状态
func (d RoleDao) UpdateRoleStatus(role *entity.SysRole) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("role_id = ?", role.RoleId).Cols("status", "update_by", "update_time").Update(role)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return update
}

// SelectRolesByUserName 查询角色组
func (d RoleDao) SelectRolesByUserName(name string) *[]entity.SysRole {
	roles := make([]entity.SysRole, 0)
	session := d.sqlSelectJoin()
	err := session.Where("r.del_flag = '0'").And("u.user_name = ?", name).Find(&roles)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &roles
}
