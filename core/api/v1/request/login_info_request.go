package request

import "cutego/pkg/base"

// LoginInfoQuery 用户get请求数据参数
type LoginInfoQuery struct {
	base.GlobalQuery
	UserName string // 筛选用户名称
}
