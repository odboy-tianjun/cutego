package base

import "time"

// GlobalModel 全局映射实体
type GlobalModel struct {
	CreateTime time.Time `json:"createTime"` // 创建时间
	CreateBy   string    `json:"createBy"`   // 创建人
	UpdateTime time.Time `json:"updateTime"` // 更新时间
	UpdateBy   string    `json:"updateBy"`   // 更新人
}

// GlobalQuery 全局Query通用条件
type GlobalQuery struct {
	BeginTime string `form:"beginTime"` // 开始时间
	EndTime   string `form:"endTime"`   // 结束时间
	PageNum   int    `form:"pageNum"`   // 当前页码
	PageSize  int    `form:"pageSize"`  // 显示条数
}
