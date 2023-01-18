package modules

import (
	router "cutego/modules/core/router"
	"cutego/pkg/filter"
	"cutego/pkg/jwt"
	"cutego/pkg/middleware"
	"cutego/pkg/middleware/logger"
	"cutego/pkg/websocket"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	ginInstance := gin.New()
	ginInstance.Use(gin.Logger())
	ginInstance.Use(gin.Recovery())
	ginInstance.Use(logger.LoggerToFile())
	ginInstance.Use(middleware.Recover)
	ginInstance.Use(jwt.JWTAuth())
	ginInstance.Use(filter.DemoHandler())
	// websocket
	ginInstance.GET("/websocket", websocket.HandleWebSocketMessage)
	// v1版本api
	v1Router := ginInstance.Group("/api/v1")
	// 加载: 模块路由
	router.LoadCoreRouter(v1Router)
	return ginInstance
}
