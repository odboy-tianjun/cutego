package dataobject

import (
	"time"
)

// SysDictData 字典数据
type SysDictData struct {
	DictCode   int64     `excel:"name=字典编码" xorm:"pk autoincr" json:"dictCode"`          // 字典ID
	DictSort   int       `excel:"name=字典排序" xorm:"int" json:"dictSort"`                  // 字典排序
	DictLabel  string    `excel:"name=字典标签" xorm:"varchar(128)" json:"dictLabel"`        // 字典标签
	DictValue  string    `excel:"name=字典键值" xorm:"varchar(128)" json:"dictValue"`        // 字典键值
	DictType   string    `excel:"name=字典类型" xorm:"varchar(128)" json:"dictType"`         // 字典类型
	IsDefault  string    `excel:"name=是否默认,format=Y=是,N=否" json:"isDefault"`             // 是否默认
	Status     string    `excel:"name=状态,format=0=正常,1=停用" xorm:"char(1)" json:"status"` // 状态 0正常1停用
	Remark     string    `xorm:"varchar(512)" json:"remark"`                             // 备注
	CreateTime time.Time `xorm:"created" json:"createTime"`                              // 创建时间
	CreateBy   string    `json:"createBy"`                                               // 创建人
	UpdateTime time.Time `xorm:"updated" json:"updateTime"`                              // 更新时间
	UpdateBy   string    `json:"updateBy"`                                               // 更新人
}

func (SysDictData) TableName() string {
	return "sys_dict_data"
}
