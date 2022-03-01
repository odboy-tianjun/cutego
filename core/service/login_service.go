package service

import (
	"cutego/core/api/v1/response"
	"cutego/pkg/common"
	"cutego/pkg/jwt"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"strings"
)

type LoginService struct {
	userService UserService
}

// Login 用户登录业务处理
func (s LoginService) Login(name string, password string) (bool, string) {
	user := s.userService.GetUserByUserName(name)
	if user == nil {
		return false, "用户不存在"
	}
	if !gotool.BcryptUtils.CompareHash(user.Password, password) {
		return false, "密码错误"
	}
	// 生成token
	token, err := jwt.CreateUserToken(s.userService.GetUserById(user.UserId))
	if err != nil {
		common.ErrorLog(err)
		return false, ""
	}
	// 数据存储到redis中
	return true, token
}

// GetCurrentUser 获取当前登录用户
func (s LoginService) GetCurrentUser(c *gin.Context) *response.UserResponse {
	token := c.Request.Header.Get("Authorization")
	str := strings.Split(token, " ")
	// parseToken 解析token包含的信息
	claims, err := jwt.ParseToken(str[1])
	if err != nil {
		common.ErrorLog(err)
	}
	info := claims.UserInfo
	return &info
}
