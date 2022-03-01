package entity

import (
	"reflect"
	"time"
)

// SysUser 用户表数据结构体
type SysUser struct {
	UserId      int64     `xorm:"pk autoincr" json:"userId"`      // 用户ID
	DeptId      int64     `json:"deptId"`                         // 部门ID
	UserName    string    `xorm:"varchar(128)" json:"userName"`   // 登录用户名
	NickName    string    `xorm:"varchar(128)" json:"nickName"`   // 用户昵称
	Email       string    `xorm:"varchar(1024)" json:"email"`     // 邮箱
	PhoneNumber string    `xorm:"varchar(11)" json:"phoneNumber"` // 手机号
	Sex         string    `xorm:"char(1)" json:"sex"`             // 性别0男1女
	Avatar      string    `xorm:"varchar(128)" json:"avatar"`     // 头像路径
	Password    string    `xorm:"varchar(128)" json:"-"`          // 密码
	Status      string    `xorm:"char(1)" json:"status"`          // 状态 0正常1停用
	DelFlag     string    `xorm:"char(1)" json:"delFlag"`         // 0正常1删除
	LoginIp     string    `xorm:"varchar(128)" json:"loginIp"`    // 登录ip
	LoginDate   time.Time `json:"loginDate"`                      // 登录时间
	CreateTime  time.Time `xorm:"created" json:"createTime"`      // 创建时间
	CreateBy    string    `json:"createBy"`                       // 创建人
	UpdateTime  time.Time `json:"updateTime"`                     // 更新时间
	UpdateBy    string    `json:"updateBy"`                       // 更新人
}

func (receiver SysUser) TableName() string {
	return "sys_user"
}

// IsAdmin 判断用户是不是管理员
func (u SysUser) IsAdmin(userId int64) bool {
	return userId > 0 && 1 == userId
}

func (a SysUser) IsEmpty() bool {
	return reflect.DeepEqual(a, SysUser{})
}
