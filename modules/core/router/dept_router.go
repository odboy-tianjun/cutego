package router

import (
	"cutego/modules/core/api/v1"
	"github.com/gin-gonic/gin"
)

// 部门接口操作
func initDeptRouter(router *gin.RouterGroup) {
	v := new(v1.DeptApi)
	group := router.Group("/dept")
	{
		// 获取部门下拉树列表
		group.GET("/treeSelect", v.DeptTreeSelect)
		// 加载对应角色部门列表树
		group.GET("/roleDeptTreeSelect/:roleId", v.RoleDeptTreeSelect)
		// 查询部门列表
		group.GET("/list", v.Find)
		// 查询部门列表（排除节点）
		group.GET("/list/exclude/:deptId", v.ExcludeChild)
		// 根据部门编号获取详细信息
		group.GET("/:deptId", v.GetInfo)
		// 添加部门
		group.POST("/create", v.Add)
		// 删除部门
		group.DELETE("/:deptId", v.Delete)
		// 修改部门
		group.PUT("/modify", v.Edit)
	}
}
