package service

import (
	"bytes"
	req2 "cutego/core/api/v1/request"
	dao2 "cutego/core/dao"
	models2 "cutego/core/entity"
	"github.com/druidcaesa/gotool"
)

type RoleService struct {
	roleDao     dao2.RoleDao
	roleMenuDao dao2.RoleMenuDao
	userRoleDao dao2.UserRoleDao
}

// FindAll 查询所有角色
func (s RoleService) FindAll(query *req2.RoleQuery) ([]*models2.SysRole, int64) {
	if query == nil {
		all := s.roleDao.SelectAll()
		return all, 0
	}
	return s.roleDao.SelectPage(query)
}

// FindRoleListByUserId 根据用户id查询角色id集合
func (s RoleService) FindRoleListByUserId(parseInt int64) *[]int64 {
	return s.roleDao.SelectRoleListByUserId(parseInt)
}

// GetRoleListByUserId 根据用户ID查询角色
func (s RoleService) GetRoleListByUserId(id int64) *[]models2.SysRole {
	return s.roleDao.GetRoleListByUserId(id)
}

// FindPage 分页查询角色数据
func (s RoleService) FindPage(query req2.RoleQuery) ([]*models2.SysRole, int64) {
	return s.roleDao.SelectPage(&query)
}

// GetRoleByRoleId 根据角色id查询角色数据
func (s RoleService) GetRoleByRoleId(id int64) *models2.SysRole {
	return s.roleDao.SelectRoleByRoleId(id)
}

// CheckRoleNameUnique 判断角色名城是否存在
func (s RoleService) CheckRoleNameUnique(role models2.SysRole) int64 {
	return s.roleDao.CheckRoleNameUnique(role)
}

// CheckRoleKeyUnique 校验角色权限是否唯一
func (s RoleService) CheckRoleKeyUnique(role models2.SysRole) int64 {
	return s.roleDao.CheckRoleKeyUnique(role)

}

// Save 添加角色数据
func (s RoleService) Save(role models2.SysRole) int64 {
	role = s.roleDao.Insert(role)
	return s.BindRoleMenu(role)
}

// 添加角色菜单关系
func (s RoleService) BindRoleMenu(role models2.SysRole) int64 {
	list := make([]models2.SysRoleMenu, 0)
	for _, id := range role.MenuIds {
		menu := models2.SysRoleMenu{
			RoleId: role.RoleId,
			MenuId: id,
		}
		list = append(list, menu)
	}
	return s.roleMenuDao.Insert(list)
}

// Edit 修改角色数据
func (s RoleService) Edit(role models2.SysRole) int64 {
	// 删除菜单关联关系
	s.roleMenuDao.Delete(role)
	s.BindRoleMenu(role)
	// 修改数据
	return s.roleDao.Update(role)
}

// Remove 删除角色
func (s RoleService) Remove(id int64) int64 {
	role := models2.SysRole{
		RoleId: id,
	}
	// 删除菜单角色关系
	s.roleMenuDao.Delete(role)
	// 删除角色
	return s.roleDao.Delete(role)
}

// CheckRoleAllowed 校验是否可以操作
func (s RoleService) CheckRoleAllowed(id int64) (bool, string) {
	if id == 1 {
		return false, "超级管理员不允许操作"
	}
	return true, ""
}

// EditRoleStatus 角色状态修改
func (s RoleService) EditRoleStatus(role *models2.SysRole) int64 {
	return s.roleDao.UpdateRoleStatus(role)
}

// DeleteAuthUser 取消授权用户
func (s RoleService) DeleteAuthUser(userRole models2.SysUserRole) int64 {
	return s.userRoleDao.DeleteAuthUser(userRole)
}

// InsertAuthUsers 批量选择用户授权
func (s RoleService) InsertAuthUsers(body req2.UserRoleBody) int64 {
	return s.userRoleDao.InsertAuthUsers(body)
}

// GetRolesByUserName 查询所属角色组
func (s RoleService) GetRolesByUserName(name string) string {
	list := s.roleDao.SelectRolesByUserName(name)
	var buffer bytes.Buffer
	var roleName string
	for _, role := range *list {
		buffer.WriteString(role.RoleName)
		buffer.WriteString(",")
	}
	s2 := buffer.String()
	if gotool.StrUtils.HasNotEmpty(s2) {
		roleName = s2[0:(len(s2) - 1)]
	}
	return roleName
}
