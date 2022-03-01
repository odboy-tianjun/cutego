package request

import "cutego/pkg/base"

type CronJobQuery struct {
	base.GlobalQuery
	JobName string `form:"jobName"`
	Status  string `form:"Status"`
}
