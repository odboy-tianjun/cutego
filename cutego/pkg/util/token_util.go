package util

import (
	"cutego/modules/core/api/v1/response"
	"cutego/modules/core/dataobject"
	"cutego/pkg/config"
	"cutego/pkg/jwt"
	"cutego/pkg/logging"
	"cutego/refs"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserUtils struct {
}

// GetUserInfo 通过jwt获取当前登录用户
func GetUserInfo(c *gin.Context) *response.UserResponse {
	token := c.Request.Header.Get("Authorization")
	s := strings.Split(token, " ")
	// parseToken 解析token包含的信息
	claims, err := jwt.ParseToken(s[1])
	if err != nil {
		logging.ErrorLog(err)
	}
	info := claims.UserInfo
	return &info
}

// CheckLockToken 校验多终端登录锁
func CheckLockToken(c *gin.Context) bool {
	if config.AppEnvConfig.Login.Single {
		// 获取redis中的token数据
		info := GetUserInfo(c)
		get, err := refs.RedisDB.GET(info.UserName)
		if err != nil {
			logging.ErrorLog(err)
			return false
		}
		token := c.Request.Header.Get(config.AppEnvConfig.Jwt.Header)
		s := strings.Split(token, " ")
		if get == s[1] {
			return true
		} else {
			return false
		}
	}
	return true
}

// CheckIsAdmin 判断是否是超级管理员
func CheckIsAdmin(user *dataobject.SysUser) bool {
	if user.UserId == 1 {
		return true
	} else {
		return false
	}
}
