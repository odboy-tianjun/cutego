package v1

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/api/v1/response"
	"cutego/modules/core/dataobject"
	"cutego/modules/core/service"
	"cutego/pkg/common"
	"cutego/pkg/excels"
	"cutego/pkg/logging"
	"cutego/pkg/page"
	"cutego/pkg/resp"
	"cutego/pkg/util"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

// UserApi 用户操作api
type UserApi struct {
	userService service.UserService
	roleService service.RoleService
	postService service.PostService
}

// Find 查询用户列表
func (a UserApi) Find(c *gin.Context) {
	query := request.UserQuery{}
	if c.BindQuery(&query) == nil {
		list, i := a.userService.FindList(query)
		success := resp.Success(page.Page{
			Size:  query.PageSize,
			Total: i,
			List:  list,
		}, "查询成功")
		c.JSON(http.StatusOK, success)
	} else {
		c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "参数错误"))
	}
}

// GetInfo 查询用户信息
func (a UserApi) GetInfo(c *gin.Context) {
	param := c.Param("userId")
	r := new(response.UserInfo)
	// 查询角色
	roleAll, _ := a.roleService.FindAll(nil)
	// 岗位所有数据
	postAll := a.postService.FindAll()
	// 判断id传入的是否为空
	if !gotool.StrUtils.HasEmpty(param) {
		parseInt, err := strconv.ParseInt(param, 10, 64)
		if err == nil {
			// 判断当前登录用户是否是admin
			m := new(dataobject.SysUser)
			if m.IsAdmin(parseInt) {
				r.Roles = roleAll
			} else {
				roles := make([]*dataobject.SysRole, 0)
				for _, role := range roleAll {
					if role.RoleId != 1 {
						roles = append(roles, role)
					}
				}
				r.Roles = roles
			}
			// 根据id获取用户数据
			r.User = a.userService.GetUserById(parseInt)
			// 根据用户ID查询岗位id集合
			r.PostIds = a.postService.FindPostListByUserId(parseInt)
			// 根据用户ID查询角色id集合
			r.RoleIds = a.roleService.FindRoleListByUserId(parseInt)
		}
	} else {
		//id为空不取管理员角色
		roles := make([]*dataobject.SysRole, 0)
		for _, role := range roleAll {
			if role.RoleId != 1 {
				roles = append(roles, role)
			}
		}
		r.Roles = roles
	}
	r.Posts = postAll
	c.JSON(http.StatusOK, resp.Success(r, "操作成功"))
}

// AuthRole 根据用户编号获取授权角色
func (a UserApi) AuthRole(c *gin.Context) {
	m := make(map[string]interface{})
	userId := c.Param("userId")
	parseInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		logging.ErrorLog(err)
		c.JSON(http.StatusInternalServerError, resp.ErrorResp(err))
	}
	user := a.userService.GetUserById(parseInt)
	// 查询角色
	roles := a.roleService.GetRoleListByUserId(parseInt)
	flag := dataobject.SysUser{}.IsAdmin(parseInt)
	if flag {
		m["roles"] = roles
	} else {
		roleList := make([]dataobject.SysRole, 0)
		for _, role := range *roles {
			if role.RoleId != 1 {
				roleList = append(roleList, role)
			}
		}
		m["roles"] = roleList
	}
	m["user"] = user
	c.JSON(http.StatusOK, resp.Success(m))
}

// Add 新增用户
func (a UserApi) Add(c *gin.Context) {
	userBody := request.UserBody{}
	if c.BindJSON(&userBody) == nil {
		// 根据用户名查询用户
		user := a.userService.GetUserByUserName(userBody.UserName)
		if user != nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败, 登录账号已存在"))
			return
		} else if a.userService.CheckPhoneNumUnique(userBody) != nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败, 手机号码已存在"))
			return
		} else if a.userService.CheckEmailUnique(userBody) != nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败, 邮箱已存在"))
			return
		}
		// 进行密码加密
		userBody.Password = gotool.BcryptUtils.Generate(userBody.Password)
		// 添加用户
		if a.userService.Save(userBody) {
			c.JSON(http.StatusOK, resp.Success(nil))
		} else {
			c.JSON(http.StatusInternalServerError, resp.ErrorResp("保存失败"))
		}
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数错误"))
	}
}

// Edit 修改用户
func (a UserApi) Edit(c *gin.Context) {
	userBody := request.UserBody{}
	if c.BindJSON(&userBody) == nil {
		if a.userService.CheckPhoneNumUnique(userBody) != nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败, 手机号码已存在"))
			return
		} else if a.userService.CheckEmailUnique(userBody) != nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败, 邮箱已存在"))
			return
		}
		// 进行用户修改操作
		if a.userService.Edit(userBody) > 0 {
			resp.OK(c)
			return
		} else {
			resp.Error(c)
			return
		}
	} else {
		resp.ParamError(c)
		return
	}
}

// Remove 删除用户
func (a UserApi) Remove(c *gin.Context) {
	param := c.Param("userId")
	userId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logging.ErrorLog(err)
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数错误"))
		return
	}
	if a.userService.Remove(userId) > 0 {
		c.JSON(http.StatusOK, resp.Success(nil))
		return
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("删除失败"))
		return
	}
}

// ResetPwd 修改重置密码
func (a UserApi) ResetPwd(c *gin.Context) {
	userBody := request.UserBody{}
	if c.BindJSON(&userBody) == nil {
		if a.userService.CheckUserAllowed(userBody) {
			c.JSON(http.StatusInternalServerError, resp.ErrorResp("不允许操作超级管理员用户"))
			return
		}
		userBody.Password = gotool.BcryptUtils.Generate(userBody.Password)
		if a.userService.ResetPwd(userBody) > 0 {
			c.JSON(http.StatusOK, resp.Success(nil))
		} else {
			c.JSON(http.StatusInternalServerError, resp.ErrorResp("重置失败"))
		}
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数错误"))
	}
}

// Export 导出excel
func (a UserApi) Export(c *gin.Context) {
	query := request.UserQuery{}
	if c.BindQuery(&query) == nil {
		items := make([]interface{}, 0)
		list, _ := a.userService.FindList(query)
		for _, userResponse := range list {
			items = append(items, *userResponse)
		}
		_, file := excels.ExportExcel(items, "用户表")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+gotool.IdUtils.IdUUIDToRan(false)+".xlsx")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("FileName", gotool.IdUtils.IdUUIDToRan(false)+".xlsx")
		file.Write(c.Writer)
	} else {
		c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "参数错误"))
	}
}

// Profile 查询个人信息
func (a UserApi) Profile(c *gin.Context) {
	m := make(map[string]interface{})
	info := util.GetUserInfo(c)
	u := a.userService.GetUserById(info.UserId)
	m["user"] = u
	// 查询所属角色组
	m["roleGroup"] = a.roleService.GetRolesByUserName(info.UserName)
	m["postGroup"] = a.postService.FindPostByUserName(info.UserName)
	resp.OK(c, m)
}

// UpdateProfile 修改个人数据
func (a UserApi) UpdateProfile(c *gin.Context) {
	user := request.UserBody{}
	if c.Bind(&user) != nil {
		resp.ParamError(c)
		return
	}
	if a.userService.CheckEmailUnique(user) != nil {
		resp.Error(c, "修改用户'"+user.UserName+"'失败, 邮箱账号已存在")
		return
	}
	if a.userService.CheckPhoneNumUnique(user) != nil {
		resp.Error(c, "修改用户'"+user.UserName+"'失败, 手机号已存在")
		return
	}
	if a.userService.EditProfile(user) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// UpdatePwd 修改个人密码
func (a UserApi) UpdatePwd(c *gin.Context) {
	oldPassword := c.Query("oldPassword")
	newPassword := c.Query("newPassword")
	info := util.GetUserInfo(c)
	name := a.userService.GetUserByUserName(info.UserName)
	hash := gotool.BcryptUtils.CompareHash(name.Password, oldPassword)
	if !hash {
		resp.Error(c, "修改密码失败, 旧密码错误")
		return
	}
	generate := gotool.BcryptUtils.Generate(oldPassword)
	compareHash := gotool.BcryptUtils.CompareHash(generate, newPassword)
	if compareHash {
		resp.Error(c, "新密码不能与旧密码相同")
		return
	}
	pwd := a.userService.EditPwd(info.UserId, gotool.BcryptUtils.Generate(newPassword))
	if pwd {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Avatar 修改头像
func (a UserApi) Avatar(c *gin.Context) {
	dirPath := common.GetDirPath("avatar")
	file, _, err := c.Request.FormFile("avatarFile")
	fileName := gotool.IdUtils.IdUUIDToRan(true) + ".jpg"
	filePath := dirPath + fileName
	fileAppend, err := gotool.FileUtils.OpenFileAppend(filePath)
	defer fileAppend.Close()
	if err != nil {
		logging.ErrorLog(err)
		resp.Error(c)
		return
	}
	_, err = io.Copy(fileAppend, file)
	if err != nil {
		logging.ErrorLog(err)
		resp.Error(c)
		return
	}
	info := util.GetUserInfo(c)
	info.Avatar = filePath
	avatar := a.userService.EditAvatar(info)
	if avatar {
		m := make(map[string]interface{})
		m["imgUrl"] = filePath
		resp.OK(c, m)
	} else {
		resp.Error(c)
	}
}

// 修改可用状态
func (a UserApi) ChangeStatus(c *gin.Context) {
	user := request.UserBody{}
	if c.Bind(&user) != nil {
		resp.ParamError(c)
		return
	}
	if user.UserId == 1 && user.Status == "1" {
		c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "不可禁用admin账号"))
		return
	}
	if a.userService.EditStatus(user) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}
