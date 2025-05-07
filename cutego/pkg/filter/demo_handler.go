package filter

import (
	"cutego/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func DemoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.AppCoreConfig.CuteGoConfig.DemoMode {
			request := inDisRequest()
			for i := 0; i < len(request); i++ {
				if c.Request.Method == http.MethodDelete || c.Request.Method == http.MethodPut || strings.Contains(c.Request.RequestURI, request[i]) {
					c.JSON(http.StatusOK, gin.H{
						"status": http.StatusInternalServerError,
						"msg":    "演示模式, 不允许操作",
					})
					c.Abort()
					return
				}
			}

		}

	}
}

// 禁用请求
func inDisRequest() []string {
	return []string{
		"/remove",
		"/profile/avatar",
		"/resetPwd",
		"/edit",
		"/insert",
		"/add",
		"/delete",
		"/export",
		"/import",
	}
}
