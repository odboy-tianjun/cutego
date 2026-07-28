package request

import "cutego/pkg/base"

type JobLogQuery struct {
	base.GlobalQuery
	JobName string `form:"jobName"`
	Status  string `form:"status"`
}