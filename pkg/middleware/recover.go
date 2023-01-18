package middleware

import (
	"cutego/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logging.ErrorLog("panic: %v\n", r)
			debug.PrintStack()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  errorToString(r),
				"data": nil,
			})
			// 终止后续接口调用, 不加的话recover捕捉到异常后, 还会继续执行接口后续的代码
			c.Abort()
		}
	}()
	// 加载完 defer recover, 继续后续接口调用
	c.Next()
}

// recover错误, 转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
