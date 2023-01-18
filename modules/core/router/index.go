package router

import "github.com/gin-gonic/gin"

func LoadCoreRouter(v1Router *gin.RouterGroup) {
	// 登录接口
	initLoginRouter(v1Router)
	// 用户路由接口
	initUserRouter(v1Router)
	// 部门路由注册
	initDeptRouter(v1Router)
	// 初始化字典数据路由
	initDictDataRouter(v1Router)
	// 注册配置路由
	initConfigRouter(v1Router)
	// 注册角色路由
	initRoleRouter(v1Router)
	// 注册菜单路由
	initMenuRouter(v1Router)
	// 注册岗位路由
	initPostRouter(v1Router)
	// 注册字典类型路由
	initDictTypeRouter(v1Router)
	// 注册公告路由
	initNoticeRouter(v1Router)
	// 注册定时任务路由
	initCronJobRouter(v1Router)
}
