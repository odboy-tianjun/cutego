package dataobject

import (
	"time"
)

type SysConfig struct {
	ConfigId    int       `excel:"name=参数主键" xorm:"pk autoincr" json:"configId"`              // 主键id
	ConfigName  string    `excel:"name=参数名称" xorm:"varchar(100)" json:"configName"`           // 参数名称
	ConfigKey   string    `excel:"name=参数键名" xorm:"varchar(100)" json:"configKey"`            // 参数建名
	ConfigValue string    `excel:"name=参数键值" xorm:"varchar(1)" json:"configValue"`            // 参数键值
	ConfigType  string    `excel:"name=系统内置,format=Y=是,N=否" xorm:"char(1)" json:"configType"` // 系统内置（Y是 N否）
	CreateBy    string    `xorm:"varchar(64)" json:"createBy"`                                // 创建人
	CreateTime  time.Time `xorm:"created" json:"createTime"`                                  // 创建时间
	UpdateBy    string    `xorm:"varchar(64)" json:"updateBy"`                                // 更新人
	UpdateTime  time.Time `xorm:"updated" json:"updateTime"`                                  // 更新时间
	Remark      string    `xorm:"varchar(500)" json:"remark"`                                 // 备注
}

func (c SysConfig) TableName() string {
	return "sys_config"
}
