package service

import (
	"cutego/modules/core/api/v1/response"
	"cutego/modules/core/dao"
	"cutego/modules/core/entity"
	"github.com/druidcaesa/gotool"
)

type PermissionService struct {
	roleDao dao.RoleDao
	menuDao dao.MenuDao
}

// GetRolePermissionByUserId 查询用户角色集合
func (s PermissionService) GetRolePermissionByUserId(user *response.UserResponse) *[]string {
	admin := entity.SysUser{}.IsAdmin(user.UserId)
	roleKeys := s.roleDao.SelectRolePermissionByUserId(user.UserId)
	if admin && roleKeys != nil {
		*roleKeys = append(*roleKeys, "admin")
	}
	duplication := gotool.StrArrayUtils.ArrayDuplication(*roleKeys)
	return &duplication
}

// GetMenuPermission 获取菜单数据权限
func (s PermissionService) GetMenuPermission(user *response.UserResponse) *[]string {
	flag := entity.SysUser{}.IsAdmin(user.UserId)
	// 查询菜单数据权限
	permission := s.menuDao.GetMenuPermission(user.UserId)
	if flag && permission != nil {
		*permission = append(*permission, "*:*:*")
	}
	var ret []string
	duplication := gotool.StrArrayUtils.ArrayDuplication(*permission)
	for i := 0; i < len(duplication); i++ {
		if (i > 0 && duplication[i-1] == duplication[i]) || len(duplication[i]) == 0 {
			continue
		}
		ret = append(ret, duplication[i])
	}
	return &ret
}
