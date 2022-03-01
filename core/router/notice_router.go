package router

import (
	"cutego/core/api/v1"
	"github.com/gin-gonic/gin"
)

func initNoticeRouter(router *gin.RouterGroup) {
	v := new(v1.NoticeApi)
	group := router.Group("/notice")
	{
		group.GET("/list", v.List)
		// 添加公告
		group.POST("/create", v.Add)
		// 删除
		group.DELETE("/:ids", v.Delete)
		// 查询
		group.GET("/:id", v.Get)
		// 修改
		group.PUT("/modify", v.Edit)
	}
}
