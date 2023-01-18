package request

import (
	"cutego/pkg/base"
	"time"
)

// UserQuery 用户get请求数据参数
type UserQuery struct {
	base.GlobalQuery
	RoleId      int64  `form:"roleId"`      // 角色id
	UserName    string `form:"userName"`    // 用户名
	Status      string `form:"status"`      // 状态
	PhoneNumber string `form:"phoneNumber"` // 手机号
	DeptId      int64  `form:"deptId"`      // 部门id
}

// UserBody 用户接收POST 或者 PUT请求参数
type UserBody struct {
	UserId      int64     `xorm:"pk autoincr" json:"userId"` // 用户ID
	DeptId      int64     `json:"deptId"`                    // 部门ID
	UserName    string    `json:"userName"`                  // 登录用户名
	NickName    string    `json:"nickName"`                  // 用户昵称
	Email       string    `json:"email"`                     // 邮箱
	PhoneNumber string    `json:"phoneNumber"`               // 手机号
	Sex         string    `json:"sex"`                       // 性别0男1女
	Avatar      string    `json:"avatar"`                    // 头像路径
	Password    string    `json:"password"`                  // 密码
	Status      string    `json:"status"`                    // 状态 0正常1停用
	DelFlag     string    `json:"delFlag"`                   // 0正常1删除
	LoginIp     string    `json:"loginIp"`                   // 登录ip
	LoginDate   time.Time `json:"loginDate"`                 // 登录时间
	CreateTime  time.Time `xorm:"created" json:"createTime"` // 创建时间
	CreateBy    string    `json:"createBy"`                  // 创建人
	UpdateTime  time.Time `json:"updateTime"`                // 更新时间
	UpdateBy    string    `json:"updateBy"`                  // 更新人
	RoleIds     []int64   `xorm:"-" json:"roleIds"`          // 角色id组
	PostIds     []int64   `xorm:"-" json:"postIds"`          // 岗位id组
}
