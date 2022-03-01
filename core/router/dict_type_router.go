package router

import (
	"cutego/core/api/v1"
	"github.com/gin-gonic/gin"
)

// 初始化字典类型路由
func initDictTypeRouter(router *gin.RouterGroup) {
	v := new(v1.DictTypeApi)
	group := router.Group("/dict/type")
	{
		// 查询字典类型分页数据
		group.GET("/list", v.List)
		// 根据id查询字典类型数据
		group.GET("/:dictTypeId", v.Get)
		// 修改字典类型
		group.PUT("/modify", v.Edit)
		// 新增字典类型
		group.POST("/create", v.Add)
		// 删除字典类型
		group.DELETE("/:dictId", v.Remove)
		// 刷新缓存
		group.DELETE("/refreshCache", v.RefreshCache)
		// 导出excel
		group.GET("/export", v.Export)
	}
}
