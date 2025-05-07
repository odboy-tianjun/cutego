package request

import "cutego/pkg/base"

type DiceDataQuery struct {
	base.GlobalQuery
	DictType  string `form:"dictType"`
	DictLabel string `form:"dictLabel"`
	Status    string `form:"status"`
}
