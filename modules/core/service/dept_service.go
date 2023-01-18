package service

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dao"
	"cutego/modules/core/entity"
)

type DeptService struct {
	deptDao dao.DeptDao
	roleDao dao.RoleDao
}

// TreeSelect 根据条件查询部门树
func (s DeptService) FindTreeSelect(query request.DeptQuery) *[]entity.SysDept {
	treeSelect := s.deptDao.SelectTree(query)
	return treeSelect
}

// FindDeptListByRoleId 根据角色ID查询部门树信息
func (s DeptService) FindDeptListByRoleId(id int64) *[]int64 {
	role := s.roleDao.SelectRoleByRoleId(id)
	return s.deptDao.SelectDeptListByRoleId(id, role.DeptCheckStrictly)
}

// FindDeptList 查询部门列表
func (s DeptService) FindDeptList(query request.DeptQuery) *[]entity.SysDept {
	return s.deptDao.GetList(query)
}

// GetDeptById 根据部门编号获取详细信息
func (s DeptService) GetDeptById(id int) *entity.SysDept {
	return s.deptDao.SelectDeptById(id)
}

// Save 添加部门数据
func (s DeptService) Save(dept entity.SysDept) int64 {
	return s.deptDao.Insert(dept)
}

// CheckDeptNameUnique 校验部门名称是否唯一
func (s DeptService) CheckDeptNameUnique(dept entity.SysDept) bool {
	if s.deptDao.CheckDeptNameUnique(dept) > 0 {
		return true
	}
	return false
}

// Remove 删除部门
func (s DeptService) Remove(id int) int64 {
	return s.deptDao.Delete(id)
}

// HasChildByDeptId 是否存在部门子节点
func (s DeptService) HasChildByDeptId(id int) int64 {
	return s.deptDao.HasChildByDeptId(id)
}

// CheckDeptExistUser 查询部门是否存在用户
func (s DeptService) CheckDeptExistUser(id int) int64 {
	return s.deptDao.CheckDeptExistUser(id)
}

// 修改部门
func (s DeptService) Edit(dept entity.SysDept) bool {
	return s.deptDao.Update(dept) > 0
}
