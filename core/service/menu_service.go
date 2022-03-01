package service

import (
	"cutego/core/api/v1/request"
	"cutego/core/api/v1/response"
	dao2 "cutego/core/dao"
	models2 "cutego/core/entity"
)

type MenuService struct {
	menuDao dao2.MenuDao
	roleDao dao2.RoleDao
}

// GetMenuTreeByUserId 根据用户ID查询菜单
func (s MenuService) GetMenuTreeByUserId(user *response.UserResponse) *[]models2.SysMenu {
	var menuList *[]models2.SysMenu
	// 判断是否是管理员
	flag := models2.SysUser{}.IsAdmin(user.UserId)
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
func (s MenuService) FindMenuList(query request.MenuQuery, info *response.UserResponse) *[]models2.SysMenu {
	if info.IsAdmin() {
		return s.menuDao.SelectMenuList(query)
	} else {
		query.UserId = info.UserId
		return s.menuDao.SelectMenuListByUserId(query)
	}
}

// GetMenuByMenuId 根据菜单ID查询信息
func (s MenuService) GetMenuByMenuId(id int) *models2.SysMenu {
	return s.menuDao.SelectMenuByMenuId(id)
}

// Save 添加菜单数据
func (s MenuService) Save(menu models2.SysMenu) int64 {
	return s.menuDao.Insert(menu)
}

// Edit 修改菜单数据
func (s MenuService) Edit(menu models2.SysMenu) int64 {
	return s.menuDao.Update(menu)
}

// Remove 删除菜单操作
func (s MenuService) Remove(id int) int64 {
	return s.menuDao.Delete(id)
}
