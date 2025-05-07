package v1

import (
	request2 "cutego/modules/core/api/v1/request"
	cache2 "cutego/modules/core/cache"
	"cutego/modules/core/dataobject"
	"cutego/modules/core/service"
	"cutego/pkg/page"
	"cutego/pkg/resp"
	"cutego/pkg/tree/tree_menu"
	"cutego/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type LoginApi struct {
	loginService      service.LoginService
	roleService       service.RoleService
	permissionService service.PermissionService
	menuService       service.MenuService
	loginInfoService  service.LoginInfoService
}

// Login 登录
func (a LoginApi) Login(c *gin.Context) {
	loginBody := request2.LoginBody{}
	if c.BindJSON(&loginBody) == nil {
		m := make(map[string]string)
		loginStatus, token := a.loginService.Login(loginBody.UserName, loginBody.Password)
		if loginStatus {
			a.handleLoginInfo(c, loginBody, loginStatus)
			// 将token存入到redis中
			cache2.SetRedisToken(loginBody.UserName, token)
			m["token"] = token
			c.JSON(http.StatusOK, resp.Success(m))
		} else {
			a.handleLoginInfo(c, loginBody, loginStatus)
			c.JSON(http.StatusOK, resp.ErrorResp(token))
		}
	} else {
		c.JSON(http.StatusOK, resp.ErrorResp(500, "参数绑定错误"))
	}
}

func (a LoginApi) handleLoginInfo(c *gin.Context, body request2.LoginBody, loginStatus bool) {
	status := util.IF(loginStatus, "1", "0")

	clientIp := a.loginInfoService.GetRequestClientIp(c)
	var location string
	if "127.0.0.1" == clientIp {
		location = "本机"
	} else {
		location = a.loginInfoService.GetLocationByIp(clientIp).Country
	}

	a.loginInfoService.Save(dataobject.SysLoginInfo{
		UserName:      body.UserName,
		IpAddr:        clientIp,
		LoginLocation: location,
		Browser:       c.GetHeader("User-Agent"),
		OS:            strings.ReplaceAll(c.GetHeader("sec-ch-ua-platform"), "\"", ""),
		Status:        status.(string),
		LoginTime:     time.Now(),
	})
}

// GetUserInfo 获取用户信息
func (a LoginApi) GetUserInfo(c *gin.Context) {
	m := make(map[string]interface{})
	user := a.loginService.GetCurrentUser(c)
	// 查询用户角色集合
	roleKeys := a.permissionService.GetRolePermissionByUserId(user)
	// 权限集合
	perms := a.permissionService.GetMenuPermission(user)
	m["roles"] = roleKeys
	m["permissions"] = perms
	m["user"] = user
	c.JSON(http.StatusOK, resp.Success(m))
}

// GetRouters 根据用户ID查询菜单
func (a LoginApi) GetRouters(c *gin.Context) {
	// 获取当前登录用户
	user := a.loginService.GetCurrentUser(c)
	menus := a.menuService.GetMenuTreeByUserId(user)
	systemMenus := tree_menu.SystemMenus{}
	systemMenus = *menus
	array := systemMenus.ConvertToINodeArray(menus)
	generateTree := tree_menu.GenerateTree(array, nil)
	c.JSON(http.StatusOK, resp.Success(generateTree))
}

// GetLoginHistory 根据用户名称查询登录记录
func (a LoginApi) GetLoginHistory(c *gin.Context) {
	// 获取当前登录用户
	user := a.loginService.GetCurrentUser(c)
	// 配置参数
	query := request2.LoginInfoQuery{}
	query.UserName = user.UserName
	// 查询
	list, i := a.loginInfoService.FindPage(query)
	success := resp.Success(page.Page{
		Size:  query.PageSize,
		Total: i,
		List:  list,
	}, "查询成功")
	c.JSON(http.StatusOK, success)
}

// Logout 退出登录
func (a LoginApi) Logout(c *gin.Context) {
	// 删除Redis缓存
	name := util.GetUserInfo(c).UserName
	cache2.RemoveRedisToken(name)
	resp.OK(c)
}
