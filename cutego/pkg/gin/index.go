package gin

import (
	"cutego/modules/core/router"
	"cutego/pkg/filter"
	"cutego/pkg/jwt"
	"cutego/pkg/logging"
	"cutego/pkg/middleware"
	"cutego/pkg/middleware/logger"
	"cutego/pkg/websocket"
	"cutego/refs"
	"github.com/gin-gonic/gin"
)

func init() {
	logging.InfoLog("CoolGin init start...")
	refs.CoolGin = gin.New()
	refs.CoolGin.Use(gin.Logger())
	refs.CoolGin.Use(gin.Recovery())
	refs.CoolGin.Use(logger.LoggerToFile())
	refs.CoolGin.Use(middleware.Recover)
	refs.CoolGin.Use(jwt.JWTAuth())
	refs.CoolGin.Use(filter.DemoHandler())
	// websocket
	refs.CoolGin.GET("/websocket", websocket.HandleWebSocketMessage)
	// v1版本api
	v1Router := refs.CoolGin.Group("/api/v1")
	// 加载: 模块路由
	router.LoadCoreRouter(v1Router)
	logging.InfoLog("CoolGin init end...")
}
