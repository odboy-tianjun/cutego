package router

import (
	"cutego/modules/core/api/v1"
	"github.com/gin-gonic/gin"
)

// 初始化字典数据路由
func initDictDataRouter(router *gin.RouterGroup) {
	v := new(v1.DictDataApi)
	group := router.Group("/dict/data")
	{
		// 根据字典类型查询字典数据信息
		group.GET("/type/:dictType", v.GetByType)
		// 查询字典数据集合
		group.GET("/list", v.List)
		// 根据id查询字典数据
		group.GET("/:dictCode", v.Get)
		// 添加字点数据
		group.POST("/create", v.Add)
		// 删除字典数据
		group.DELETE("/:dictCode", v.Delete)
		// 导出字典数据
		group.GET("/export", v.Export)
		// 编辑字典数据
		group.PUT("/modify", v.Edit)
	}
}
