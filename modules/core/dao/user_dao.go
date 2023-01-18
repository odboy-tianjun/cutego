package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/api/v1/response"
	"cutego/modules/core/entity"
	"cutego/pkg/common"
	"cutego/pkg/page"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
	"time"
)

type UserDao struct {
}

// 查询公共sql
func (d UserDao) sqlSelectJoin() *xorm.Session {
	return SqlDB.NewSession().Table([]string{"sys_user", "u"}).
		Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id").
		Join("LEFT", []string{"sys_user_role", "ur"}, "u.user_id = ur.user_id").
		Join("LEFT", []string{"sys_role", "r"}, "r.role_id = ur.role_id")
}

// SelectPage 查询用户集合
func (d UserDao) SelectPage(query request.UserQuery) ([]*response.UserResponse, int64) {
	resp := make([]*response.UserResponse, 0)
	sql := d.sqlSelectJoin()
	if !gotool.StrUtils.HasEmpty(query.UserName) {
		sql.And("u.user_name like concat('%',?,'%')", query.UserName)
	}
	if !gotool.StrUtils.HasEmpty(query.Status) {
		sql.And("i.status = ?", query.Status)
	}
	if !gotool.StrUtils.HasEmpty(query.PhoneNumber) {
		sql.And("u.phone_number like concat('%',?,'%')", query.PhoneNumber)
	}
	if !gotool.StrUtils.HasEmpty(query.BeginTime) {
		sql.And("date_format(u.create_time,'%y%m%d') >= date_format(?,'%y%m%d')", query.BeginTime)
	}
	if !gotool.StrUtils.HasEmpty(query.EndTime) {
		sql.And("date_format(u.create_time,'%y%m%d') <= date_format(?,'%y%m%d')", query.EndTime)
	}
	if query.DeptId > 0 {
		sql.And("u.dept_id = ? OR u.dept_id in ( SELECT t.dept_id FROM sys_dept t WHERE find_in_set(?, ancestors))", query.DeptId, query.DeptId)
	}
	total, _ := page.GetTotal(sql.Clone())
	err := sql.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&resp)
	if err != nil {
		common.ErrorLog(err)
		return nil, 0
	}
	return resp, total
}

// GetUserById 根据id查询用户数据
func (d UserDao) GetUserById(userId int64) *response.UserResponse {
	var resp response.UserResponse
	get, err := d.sqlSelectJoin().Where("u.user_id = ?", userId).Get(&resp)
	if err != nil {
		common.ErrorLog(err)
	}
	if !get {
		return nil
	}
	return &resp
}

// GetUserByUserName 根据用户名查询用户数据
func (d UserDao) GetUserByUserName(user entity.SysUser) *entity.SysUser {
	i, err := SqlDB.Get(&user)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	if i {
		return &user
	}
	return nil
}

// CheckEmailUnique 校验邮箱是否存在
func (d UserDao) CheckEmailUnique(user request.UserBody) *entity.SysUser {
	sysUser := entity.SysUser{}
	session := SqlDB.NewSession().Table("sys_user")
	session.Cols("user_id", "email")
	session.Where("email = ?", user.Email)
	if user.UserId > 0 {
		session.And("user_id != ?", user.UserId)
	}
	get, _ := session.Limit(1).Get(&sysUser)
	if !get {
		return nil
	}
	return &sysUser
}

// CheckPhoneNumUnique 校验手机号是否存在
func (d UserDao) CheckPhoneNumUnique(body request.UserBody) *entity.SysUser {
	sysUser := entity.SysUser{}
	session := SqlDB.NewSession().Table("sys_user")
	session.Cols("user_id", "phone_num")
	session.Where("phone_num = ?", body.PhoneNumber)
	if body.UserId > 0 {
		session.And("user_id != ?", body.UserId)
	}
	get, _ := session.Limit(1).Get(&sysUser)
	if !get {
		return nil
	}
	return &sysUser
}

// Insert 添加用户
func (d UserDao) Insert(body request.UserBody) *request.UserBody {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Table("sys_user").Insert(&body)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
	return &body
}

// Update 修改用户数据
func (d UserDao) Update(body request.UserBody) int64 {
	session := SqlDB.NewSession().Table("sys_user")
	session.Begin()
	_, err := session.Where("user_id = ?", body.UserId).Update(&body)
	if err != nil {
		session.Rollback()
		common.ErrorLog(err)
		return 0
	}
	session.Commit()
	return 1
}

// Delete 根据id删除用户数据
func (d UserDao) Delete(id int64) int64 {
	user := entity.SysUser{
		UserId: id,
	}
	session := SqlDB.NewSession().Table("sys_user")
	session.Begin()
	i, err := session.Delete(&user)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
	return i
}

// ResetPwd 修改用户密码数据库操作
func (d UserDao) ResetPwd(body request.UserBody) int64 {
	user := entity.SysUser{
		UserId:   body.UserId,
		Password: body.Password,
	}
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Where("user_id = ?", user.UserId).Cols("password").Update(&user)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return 1
}

// GetAllocatedList 查询未分配用户角色列表
func (d UserDao) GetAllocatedList(query request.UserQuery) ([]*response.UserResponse, int64) {
	resp := make([]*response.UserResponse, 0)
	session := SqlDB.NewSession()
	session.Table([]string{"sys_user", "u"}).Distinct("u.user_id", "u.dept_id", "u.user_name", "u.nick_name", "u.email", "u.phone_number", "u.status", "u.create_time").
		Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id").
		Join("LEFT", []string{"sys_user_role", "ur"}, "u.user_id = ur.user_id").
		Join("LEFT", []string{"sys_role", "r"}, "r.role_id = ur.role_id").Where("u.del_flag = '0'").And("r.role_id = ?", query.RoleId)
	if gotool.StrUtils.HasNotEmpty(query.UserName) {
		session.And("u.user_name like concat('%', ?, '%')", query.UserName)
	}
	if gotool.StrUtils.HasNotEmpty(query.PhoneNumber) {
		session.And("u.phone_number like concat('%', ?, '%')", query.PhoneNumber)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&resp)
	if err != nil {
		common.ErrorLog(err)
		return nil, 0
	}
	return resp, total
}

// GetUnallocatedList 查询未分配用户角色列表
func (d UserDao) GetUnallocatedList(query request.UserQuery) ([]*response.UserResponse, int64) {
	resp := make([]*response.UserResponse, 0)
	session := SqlDB.NewSession()
	session.Table([]string{"sys_user", "u"}).Distinct("u.user_id", "u.dept_id", "u.user_name", "u.nick_name", "u.email", "u.phone_number", "u.status", "u.create_time").
		Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id").
		Join("LEFT", []string{"sys_user_role", "ur"}, "u.user_id = ur.user_id").
		Join("LEFT", []string{"sys_role", "r"}, "r.role_id = ur.role_id").Where("u.del_flag = '0'").And("r.role_id = ? or r.role_id IS NULL", query.RoleId).
		And("u.user_id not in (select u.user_id from sys_user u inner join sys_user_role ur on u.user_id = ur.user_id and ur.role_id = ?)", query.RoleId)
	if gotool.StrUtils.HasNotEmpty(query.UserName) {
		session.And("u.user_name like concat('%', ?, '%')", query.UserName)
	}
	if gotool.StrUtils.HasNotEmpty(query.PhoneNumber) {
		session.And("u.phone_number like concat('%', ?, '%')", query.PhoneNumber)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&resp)
	if err != nil {
		common.ErrorLog(err)
		return nil, 0
	}
	return resp, total
}

// UpdatePwd 修改密码
func (d UserDao) UpdatePwd(id int64, hash string) int64 {
	user := entity.SysUser{}
	user.UserId = id
	user.Password = hash
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Cols("password").Where("user_id = ?", id).Update(&user)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return update
}

// UpdateAvatar 修改头像
func (d UserDao) UpdateAvatar(info *response.UserResponse) int64 {
	user := entity.SysUser{
		Avatar:     info.Avatar,
		UserId:     info.UserId,
		UpdateBy:   info.UserName,
		UpdateTime: time.Now(),
	}
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Cols("avatar", "update_by", "update_time").Where("user_id = ?", user.UserId).Update(&user)
	if err != nil {
		session.Rollback()
		common.ErrorLog(err)
		return 0
	}
	session.Commit()
	return update
}

func (d UserDao) UpdateStatus(info request.UserBody) int64 {
	user := entity.SysUser{
		UserId:     info.UserId,
		Status:     info.Status,
		UpdateBy:   info.UserName,
		UpdateTime: time.Now(),
	}
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Cols("status", "update_by", "update_time").Where("user_id = ?", user.UserId).Update(&user)
	if err != nil {
		session.Rollback()
		common.ErrorLog(err)
		return 0
	}
	session.Commit()
	return update
}
