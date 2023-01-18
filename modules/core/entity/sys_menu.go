package entity

import (
	"strings"
	"time"
)

// SysMenu 菜单结构体
type SysMenu struct {
	MenuId     int       `xorm:"pk autoincr" json:"menuId"`     // 主键Id
	ParentId   int       `json:"parentId"`                      // 父Id
	MenuName   string    `xorm:"varchar(128)" json:"menuName"`  // 菜单名称
	OrderNum   int       `xorm:"int" json:"orderNum"`           // 显示顺序
	Path       string    `xorm:"varchar(200)" json:"path"`      // 请求路径
	MenuType   string    `xorm:"char(1)" json:"menuType"`       // 菜单类型 （M目录 C菜单 F按钮）
	Visible    string    `xorm:"char(1)" json:"visible"`        // 菜单状态 （0显示 1隐藏）
	IsFrame    int       `json:"isFrame"`                       // 是否为外链（0是 1否）
	IsCache    int       `json:"isCache"`                       // 是否缓存（0缓存 1不缓存）
	Perms      string    `xorm:"varchar(100)" json:"perms"`     // 权限标识
	Icon       string    `xorm:"varchar(100)" json:"icon"`      // 图标
	Remark     string    `xorm:"varchar(512)" json:"remark"`    // 备注
	CreateTime time.Time `xorm:"created" json:"createTime"`     // 创建时间
	CreateBy   string    `json:"createBy"`                      // 创建人
	UpdateTime time.Time `xorm:"updated" json:"updateTime"`     // 更新时间
	UpdateBy   string    `json:"updateBy"`                      // 更新人
	Status     string    `xorm:"char(1)" json:"status"`         // 菜单状态（0正常 1停用）
	Component  string    `xorm:"varchar(255)" json:"component"` // 组件路径
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
func (s SysMenu) GetPath() string {
	return s.Path
}
func (s SysMenu) GetName() string {
	return strings.Title(s.Path)
}
func (s SysMenu) GetMenuId() int {
	return s.MenuId
}
func (s SysMenu) GetParentId() int {
	return s.ParentId
}
func (s SysMenu) GetData() interface{} {
	return s
}
func (s SysMenu) IsRoot() bool {
	// 这里通过FatherId等于0 或者 FatherId等于自身Id表示顶层根节点
	return s.ParentId == 0 || s.ParentId == s.MenuId
}

func (s SysMenu) GetId() int {
	return s.MenuId
}

func (s SysMenu) GetLabel() string {
	return s.MenuName
}
