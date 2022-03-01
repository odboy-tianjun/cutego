package v1

import (
	"cutego/core/api/v1/request"
	"cutego/core/entity"
	"cutego/core/service"
	"cutego/pkg/resp"
	"cutego/pkg/tree/tree_menu"
	"cutego/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type MenuApi struct {
	menuService service.MenuService
}

// List 查询菜单数据
func (a MenuApi) List(c *gin.Context) {
	// 获取当前登录用户
	info := util.GetUserInfo(c)
	// 获取参数
	query := request.MenuQuery{}
	if c.Bind(&query) != nil {
		resp.Error(c)
		return
	}
	resp.OK(c, a.menuService.FindMenuList(query, info))
}

// GetInfo 根据id查询菜单详情
func (a MenuApi) GetInfo(c *gin.Context) {
	param := c.Param("menuId")
	menuId, err := strconv.Atoi(param)
	if err != nil {
		resp.ParamError(c, "参数绑定错误")
		return
	}
	resp.OK(c, a.menuService.GetMenuByMenuId(menuId))
}

// RoleMenuTreeSelect 加载对应角色菜单列表树
func (a MenuApi) RoleMenuTreeSelect(c *gin.Context) {
	m := make(map[string]interface{})
	param := c.Param("roleId")
	roleId, _ := strconv.ParseInt(param, 10, 64)
	// 获取当前登录用户
	info := util.GetUserInfo(c)
	menuList := a.menuService.GetMenuTreeByUserId(info)
	menus := tree_menu.SystemMenus{}
	tree := menus.GetTree(menuList)
	ids := a.menuService.FindMenuListByRoleId(roleId)
	m["checkedKeys"] = ids
	m["menus"] = tree
	c.JSON(http.StatusOK, resp.Success(m))
}

// TreeSelect 获取菜单下拉树列表
func (a MenuApi) TreeSelect(c *gin.Context) {
	info := util.GetUserInfo(c)
	menus := a.menuService.GetMenuTreeByUserId(info)
	systemMenus := tree_menu.SystemMenus{}
	tree := systemMenus.GetTree(menus)
	c.JSON(http.StatusOK, resp.Success(tree))
}

// Add 添加菜单数据
func (a MenuApi) Add(c *gin.Context) {
	menu := entity.SysMenu{}
	if c.Bind(&menu) != nil {
		resp.ParamError(c, "参数绑定异常")
		return
	}
	if a.menuService.Save(menu) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Edit 修改菜单数据
func (a MenuApi) Edit(c *gin.Context) {
	menu := entity.SysMenu{}
	if c.Bind(&menu) != nil {
		resp.ParamError(c)
		return
	}
	menu.UpdateBy = util.GetUserInfo(c).UserName
	menu.UpdateTime = time.Now()
	if a.menuService.Edit(menu) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Delete 删除菜单
func (a MenuApi) Delete(c *gin.Context) {
	param := c.Param("menuId")
	menuId, err := strconv.Atoi(param)
	if err != nil {
		resp.ParamError(c)
		return
	}
	if a.menuService.Remove(menuId) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}
