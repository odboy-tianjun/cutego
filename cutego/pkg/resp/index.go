package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 数据返回结构体
type Response struct {
	Status int         `json:"status"` // 返回状态值
	Msg    string      `json:"msg"`    // 返回的提示语
	Data   interface{} `json:"data"`   // 返回数据
}

// Success 正确返回
func Success(data interface{}, msg ...string) *Response {
	response := Response{
		Status: http.StatusOK,
		Data:   data,
		Msg:    "操作成功",
	}
	if len(msg) > 0 {
		response.Msg = msg[0]
	}
	return &response
}

// ErrorResp 错误返回
func ErrorResp(data ...interface{}) *Response {
	response := Response{
		Status: http.StatusInternalServerError,
		Msg:    "操作失败",
		Data:   nil,
	}
	for _, value := range data {
		switch value.(type) {
		case string:
			response.Msg = value.(string)
		case int:
			response.Status = value.(int)
		case interface{}:
			response.Data = value.(interface{})
		}
	}
	return &response
}

func Error(c *gin.Context, data ...interface{}) {
	response := Response{
		Status: http.StatusInternalServerError,
		Msg:    "操作失败",
		Data:   nil,
	}
	for _, value := range data {
		switch value.(type) {
		case string:
			response.Msg = value.(string)
		case int:
			response.Status = value.(int)
		case interface{}:
			response.Data = value.(interface{})
		}
	}
	c.JSON(http.StatusOK, response)
	return
}
func ParamError(c *gin.Context, data ...interface{}) {
	response := Response{
		Status: http.StatusInternalServerError,
		Msg:    "参数绑定异常",
		Data:   nil,
	}
	for _, value := range data {
		switch value.(type) {
		case string:
			response.Msg = value.(string)
		case int:
			response.Status = value.(int)
		case interface{}:
			response.Data = value.(interface{})
		}
	}
	c.JSON(http.StatusBadRequest, response)
	return
}
func OK(c *gin.Context, data ...interface{}) {
	response := Response{
		Status: http.StatusOK,
		Msg:    "操作成功",
		Data:   nil,
	}
	for _, datum := range data {
		switch datum.(type) {
		case string:
			response.Msg = datum.(string)
		case interface{}:
			response.Data = datum.(interface{})
		}
	}
	c.JSON(http.StatusOK, response)
	return
}
