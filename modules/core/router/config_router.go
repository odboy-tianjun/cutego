package router

import (
	"cutego/modules/core/api/v1"
	"github.com/gin-gonic/gin"
)

func initConfigRouter(router *gin.RouterGroup) {
	v := new(v1.ConfigApi)
	group := router.Group("/config")
	{
		// 根据参数键名查询参数值
		group.GET("", v.GetConfigValue)
		// 查询设置列表
		group.GET("/list", v.List)
		// 添加
		group.POST("/create", v.Add)
		// 查询
		group.GET("/:configId", v.Get)
		// 修改
		group.PUT("/modify", v.Edit)
		// 批量删除
		group.DELETE("/:ids", v.Delete)
		// 刷新缓存
		group.DELETE("/refreshCache", v.RefreshCache)
		// 导出数据
		group.GET("/export", v.Export)
	}
}
