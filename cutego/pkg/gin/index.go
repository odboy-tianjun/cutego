package gin

import (
	"cutego/modules/core/router"
	"cutego/pkg/config"
	"cutego/pkg/filter"
	"cutego/pkg/jwt"
	"cutego/pkg/logger"
	"cutego/pkg/middleware"
	"cutego/pkg/util"
	"cutego/pkg/websocket"
	"cutego/refs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	refs.CoolGin = gin.New()
	refs.CoolGin.Use(gin.Logger())
	refs.CoolGin.Use(gin.Recovery())
	refs.CoolGin.Use(middleware.Recover)
	refs.CoolGin.Use(jwt.JWTAuth())
	refs.CoolGin.Use(filter.DemoHandler())
	// websocket
	refs.CoolGin.GET("/websocket", websocket.HandleWebSocketMessage)
	// v1版本api
	v1Router := refs.CoolGin.Group("/api/v1")
	// 加载: 模块路由
	router.LoadCoreRouter(v1Router)
	gin.SetMode(util.IF(config.AppEnvConfig.Server.RunMode == "", "debug", config.AppEnvConfig.Server.RunMode).(string))
	err := refs.CoolGin.Run(fmt.Sprintf(":%d", config.AppEnvConfig.Server.Port))
	if err != nil {
		logger.SugaredLogger.Fatalf("Start server: %+v", err)
	}
}
