package router

import (
	"cutego/modules/core/api/v1"
	"github.com/gin-gonic/gin"
)

func initMenuRouter(router *gin.RouterGroup) {
	v := new(v1.MenuApi)
	vg := router.Group("/menu")
	{
		// 查询菜单数据
		vg.GET("/:menuId", v.GetInfo)
		// 查询菜单列表
		vg.GET("/list", v.List)
		// 加载对应角色菜单列表树
		vg.GET("/roleMenuTreeSelect/:roleId", v.RoleMenuTreeSelect)
		// 获取菜单下拉树列表
		vg.GET("/treeSelect", v.TreeSelect)
		// 添加菜单数据
		vg.POST("/create", v.Add)
		// 修改菜单数据
		vg.PUT("/modify", v.Edit)
		// 删除菜单数据
		vg.DELETE("/:menuId", v.Delete)
	}
}
