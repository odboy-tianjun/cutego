package response

// DictDataResponse 字典数据实体返回结构体
type DictDataResponse struct {
	DictCode  int64  `excel:"name=字典编码" xorm:"pk autoincr" json:"dictCode"`   // 字典ID
	DictSort  int    `excel:"name=字典排序" xorm:"int" json:"dictSort"`           // 字典排序
	DictLabel string `excel:"name=字典标签" xorm:"varchar(128)" json:"dictLabel"` // 字典标签
	DictValue string `excel:"name=字典键值" xorm:"varchar(128)" json:"dictValue"` // 字典键值
	DictType  string `excel:"name=字典类型" xorm:"varchar(128)" json:"dictType"`  // 字典类型
	IsDefault string `excel:"name=是否默认,format=Y=是,N=否" json:"isDefault"`      // 是否默认
}
