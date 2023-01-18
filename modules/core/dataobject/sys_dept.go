package dataobject

import "time"

// SysDept 部门结构体
type SysDept struct {
	DeptId     int       `xorm:"pk autoincr" json:"deptId"`
	Ancestors  string    `xorm:"varchar(50)" json:"ancestors"`
	DeptName   string    `xorm:"varchar(128)" json:"deptName"`
	OrderNum   int       `json:"orderNum"`
	Leader     string    `xorm:"varchar(20)" json:"leader"`
	ParentId   int       `json:"parentId"`
	Phone      string    `xorm:"varchar(11)" json:"phone"`
	Status     string    `xorm:"char(1)" json:"status"`
	Email      string    `json:"email"`
	DelFlag    string    `xorm:"char(1) default('0')" json:"delFlag"`
	CreateTime time.Time `xorm:"created" json:"createTime"` // 创建时间
	CreateBy   string    `json:"createBy"`                  // 创建人
	UpdateTime time.Time `xorm:"updated" json:"updateTime"` // 更新时间
	UpdateBy   string    `json:"updateBy"`                  // 更新人
}

func (d SysDept) GetLabel() string {
	return d.DeptName
}

func (d SysDept) GetId() int {
	return d.DeptId
}

func (d SysDept) GetParentId() int {
	return d.ParentId
}

func (d SysDept) GetData() interface{} {
	return d
}

func (d SysDept) IsRoot() bool {
	return d.ParentId == 0 || d.ParentId == d.DeptId
}

func (SysDept) TableName() string {
	return "sys_dept"
}
