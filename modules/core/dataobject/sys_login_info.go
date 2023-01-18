package dataobject

import (
	"time"
)

type SysLoginInfo struct {
	Id            int       `excel:"name=主键id" xorm:"pk autoincr" json:"id"`                  // 主键id
	UserName      string    `excel:"name=用户账号" xorm:"varchar(50)" json:"userName"`            // 用户账号
	IpAddr        string    `excel:"name=登录IP地址" xorm:"varchar(128)" json:"ipAddr"`           // 登录IP地址
	LoginLocation string    `excel:"name=登录地点" xorm:"varchar(255)" json:"loginLocation"`      // 登录地点
	Browser       string    `excel:"name=浏览器类型" xorm:"varchar(255)" json:"browser"`           // 浏览器类型
	OS            string    `excel:"name=操作系统" xorm:"os varchar(50)" json:"os"`               // 操作系统
	Status        string    `excel:"name=登录状态,format=1=成功,0=失败" xorm:"char(1)" json:"status"` // 登录状态（1成功 0失败）
	Msg           string    `excel:"name=提示消息" xorm:"varchar(255)" json:"msg"`                // 提示消息
	LoginTime     time.Time `excel:"name=登录时间" xorm:"login_time" json:"loginTime"`            // 登录时间
}

func (c SysLoginInfo) TableName() string {
	return "sys_login_info"
}
