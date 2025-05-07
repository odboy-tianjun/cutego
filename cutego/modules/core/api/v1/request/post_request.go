package request

import "cutego/pkg/base"

type PostQuery struct {
	base.GlobalQuery
	PostCode string `form:"postCode"`
	Status   string `form:"status"`
	PostName string `form:"postName"`
}
