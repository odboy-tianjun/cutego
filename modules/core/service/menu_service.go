package service

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/api/v1/response"
	"cutego/modules/core/dao"
	"cutego/modules/core/entity"
)

type MenuService struct {
	menuDao dao.MenuDao
	roleDao dao.RoleDao
}

// GetMenuTreeByUserId 根据用户ID查询菜单
func (s MenuService) GetMenuTreeByUserId(user *response.UserResponse) *[]entity.SysMenu {
	var menuList *[]entity.SysMenu
	// 判断是否是管理员
	flag := entity.SysUser{}.IsAdmin(user.UserId)
	if flag {
		menuList = s.menuDao.GetMenuAll()
	} else {
		menuList = s.menuDao.GetMenuByUserId(user.UserId)
	}
	return menuList
}

// FindMenuListByRoleId 根据角色ID查询菜单树信息
func (s MenuService) FindMenuListByRoleId(id int64) *[]int64 {
	role := s.roleDao.SelectRoleByRoleId(id)
	return s.menuDao.SelectMenuByRoleId(id, role.MenuCheckStrictly)
}

// GetMenuList 获取菜单列表
func (s MenuService) FindMenuList(query request.MenuQuery, info *response.UserResponse) *[]entity.SysMenu {
	if info.IsAdmin() {
		return s.menuDao.SelectMenuList(query)
	} else {
		query.UserId = info.UserId
		return s.menuDao.SelectMenuListByUserId(query)
	}
}

// GetMenuByMenuId 根据菜单ID查询信息
func (s MenuService) GetMenuByMenuId(id int) *entity.SysMenu {
	return s.menuDao.SelectMenuByMenuId(id)
}

// Save 添加菜单数据
func (s MenuService) Save(menu entity.SysMenu) int64 {
	return s.menuDao.Insert(menu)
}

// Edit 修改菜单数据
func (s MenuService) Edit(menu entity.SysMenu) int64 {
	return s.menuDao.Update(menu)
}

// Remove 删除菜单操作
func (s MenuService) Remove(id int) int64 {
	return s.menuDao.Delete(id)
}
