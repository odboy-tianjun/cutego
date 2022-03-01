package router

import (
	"cutego/core/api/v1"
	"github.com/gin-gonic/gin"
)

// 初始化定时任务路由
func initCronJobRouter(router *gin.RouterGroup) {
	v := new(v1.CronJobApi)
	group := router.Group("/cronJob")
	{
		// 查询定时任务分页数据
		group.GET("/list", v.List)
		// 修改定时任务
		group.PUT("/modify", v.Edit)
		// 新增定时任务
		group.POST("/create", v.Add)
		// 删除定时任务
		group.DELETE("/:jobId", v.Remove)
		// 改变定时任务状态
		group.DELETE("/changeStatus", v.ChangeStatus)
	}
}
