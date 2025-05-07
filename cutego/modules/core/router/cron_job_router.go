package router

import (
	"cutego/modules/core/api/v1"
	"github.com/gin-gonic/gin"
)

// 初始化定时任务路由
func initCronJobRouter(router *gin.RouterGroup) {
	v := new(v1.CronJobApi)
	group := router.Group("/monitor/cronJob")
	{
		// 查询定时任务分页数据
		group.GET("/list", v.List)
		// 查询定时任务分页数据
		group.GET("/:jobId", v.GetById)
		// 修改定时任务
		group.PUT("/modify", v.Edit)
		// 新增定时任务
		group.POST("/create", v.Add)
		// 删除定时任务
		group.DELETE("/:jobId/:funcAlias", v.Remove)
		// 改变定时任务状态
		group.PUT("/changeStatus", v.ChangeStatus)
	}
}
