package service

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dao"
	"cutego/modules/core/dataobject"
	"cutego/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/yinheli/qqwry"
	"net"
	"os"
)

type LoginInfoService struct {
	loginInfoDao dao.LoginInfoDao
}

// FindPage 分页查询数据
func (s LoginInfoService) FindPage(query request.LoginInfoQuery) (*[]dataobject.SysLoginInfo, int64) {
	return s.loginInfoDao.SelectPage(query)
}

// Save 添加登录记录业务逻辑
func (s LoginInfoService) Save(body dataobject.SysLoginInfo) bool {
	// 添加登录记录数据库操作
	user := s.loginInfoDao.Insert(body)
	if user != nil {
		return true
	}
	return false
}

// GetRequestClientIp 获取请求客户端的ip
func (s LoginInfoService) GetRequestClientIp(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

// GetLocationByIp 纯真数据库获取ip地址
// @return {"Ip": "IP地址","Country": "国家","City": "城市"}
func (s LoginInfoService) GetLocationByIp(ipAddr string) *qqwry.QQwry {
	address := net.ParseIP(ipAddr)
	if ipAddr == "" || address == nil {
		panic("参数ipAddr是空的")
	} else {
		dir, err := os.Getwd()
		if err != nil {
			panic("无法获取当前路径, " + err.Error())
		}
		q := qqwry.NewQQwry(dir + "/" + config.BaseConfigDirPath + "/scanip/qqwry.dat")
		q.Find(ipAddr)
		return q
	}
}
