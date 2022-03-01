package service

import (
	"cutego/core/api/v1/request"
	"cutego/core/api/v1/response"
	dao2 "cutego/core/dao"
	models2 "cutego/core/entity"
)

// UserService 用户操作业务逻辑
type UserService struct {
	userDao     dao2.UserDao
	userPostDao dao2.UserPostDao
	userRoleDao dao2.UserRoleDao
}

// FindList 查询用户集合业务方法
func (s UserService) FindList(query request.UserQuery) ([]*response.UserResponse, int64) {
	return s.userDao.SelectPage(query)
}

// GetUserById 根据id查询用户数据
func (s UserService) GetUserById(parseInt int64) *response.UserResponse {
	return s.userDao.GetUserById(parseInt)
}

// GetUserByUserName 根据用户名查询用户
func (s UserService) GetUserByUserName(name string) *models2.SysUser {
	user := models2.SysUser{}
	user.UserName = name
	return s.userDao.GetUserByUserName(user)
}

// CheckEmailUnique 校验邮箱是否存在
func (s UserService) CheckEmailUnique(user request.UserBody) *models2.SysUser {
	return s.userDao.CheckEmailUnique(user)
}

// CheckPhoneNumUnique 校验手机号是否存在
func (s UserService) CheckPhoneNumUnique(body request.UserBody) *models2.SysUser {
	return s.userDao.CheckPhoneNumUnique(body)
}

// Insert 添加用户业务逻辑
func (s UserService) Save(body request.UserBody) bool {
	// 添加用户数据库操作
	user := s.userDao.Insert(body)
	if user != nil {
		s.BindUserPost(user)
		s.BindUserRole(user)
		return true
	}
	return false
}

// 新增用户岗位信息
func (s UserService) BindUserPost(user *request.UserBody) {
	postIds := user.PostIds
	if len(postIds) > 0 {
		sysUserPosts := make([]models2.SysUserPost, 0)
		for i := 0; i < len(postIds); i++ {
			m := models2.SysUserPost{
				UserId: user.UserId,
				PostId: postIds[i],
			}
			sysUserPosts = append(sysUserPosts, m)
		}
		s.userPostDao.BatchInsert(sysUserPosts)
	}
}

// 新增用户角色信息
func (s UserService) BindUserRole(user *request.UserBody) {
	roleIds := user.RoleIds
	if len(roleIds) > 0 {
		roles := make([]models2.SysUserRole, 0)
		for i := 0; i < len(roleIds); i++ {
			role := models2.SysUserRole{
				RoleId: roleIds[i],
				UserId: user.UserId,
			}
			roles = append(roles, role)
		}
		s.userRoleDao.BatchInsert(roles)
	}
}

// Edit 修改用户数据
func (s UserService) Edit(body request.UserBody) int64 {
	// 删除原有用户和角色关系
	s.userRoleDao.Delete(body.UserId)
	// 重新添加用具角色关系
	s.BindUserRole(&body)
	// 删除原有用户岗位关系
	s.userPostDao.Delete(body.UserId)
	// 添加新的用户岗位关系
	s.BindUserPost(&body)
	// 修改用户数据
	return s.userDao.Update(body)
}

// Remove 根据用户id删除用户相关数据
func (s UserService) Remove(id int64) int64 {
	// 删除原有用户和角色关系
	s.userRoleDao.Delete(id)
	// 删除原有用户岗位关系
	s.userPostDao.Delete(id)
	// 删除用户数据
	return s.userDao.Delete(id)
}

// CheckUserAllowed 校验是否可以修改用户密码
func (s UserService) CheckUserAllowed(body request.UserBody) bool {
	user := models2.SysUser{}
	return user.IsAdmin(body.UserId)
}

// ResetPwd 修改用户密码
func (s UserService) ResetPwd(body request.UserBody) int64 {
	return s.userDao.ResetPwd(body)
}

// GetAllocatedList 查询未分配用户角色列表
func (s UserService) GetAllocatedList(query request.UserQuery) ([]*response.UserResponse, int64) {
	return s.userDao.GetAllocatedList(query)
}

// GetUnallocatedList 查询未分配用户角色列表
func (s UserService) GetUnallocatedList(query request.UserQuery) ([]*response.UserResponse, int64) {
	return s.userDao.GetUnallocatedList(query)
}

// EditProfile 修改数据
func (s UserService) EditProfile(user request.UserBody) int64 {
	return s.userDao.Update(user)
}

// EditPwd 修改密码
func (s UserService) EditPwd(id int64, hash string) bool {
	return s.userDao.UpdatePwd(id, hash) > 0
}

// EditAvatar 修改头像
func (s UserService) EditAvatar(info *response.UserResponse) bool {
	return s.userDao.UpdateAvatar(info) > 0
}

// EditStatus 修改可用状态
func (s UserService) EditStatus(info request.UserBody) bool {
	return s.userDao.UpdateStatus(info) > 0
}
