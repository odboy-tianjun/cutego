package router

import (
	"cutego/core/api/v1"
	"github.com/gin-gonic/gin"
)

// 初始化岗位路由
func initPostRouter(router *gin.RouterGroup) {
	v := new(v1.PostApi)
	group := router.Group("/post")
	{
		// 查询岗位数据
		group.GET("/list", v.List)
		// 添加岗位
		group.POST("/create", v.Add)
		// 查询岗位详情
		group.GET("/:postId", v.Get)
		// 删除岗位数据
		group.DELETE("/:postId", v.Delete)
		// 修改岗位数据接口
		group.PUT("/modify", v.Edit)
		// 导出excel
		group.GET("/export", v.Export)
	}
}
