package v1

import (
	"cutego/core/api/v1/request"
	"cutego/core/api/v1/response"
	"cutego/core/entity"
	"cutego/core/service"
	"cutego/pkg/common"
	"cutego/pkg/cronjob"
	"cutego/pkg/resp"
	"cutego/pkg/util"
	"github.com/gin-gonic/gin"
)

type CronJobApi struct {
	cronJobService service.CronJobService
}

// List 获取定时任务数据
func (a CronJobApi) List(c *gin.Context) {
	query := request.CronJobQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	find, i := a.cronJobService.FindPage(query)

	var result []interface{}
	for _, item := range find {
		// 领域对象转换
		result = append(result, util.DeepCopy(item, &response.CronJobResponse{}))
	}
	resp.OK(c, util.NewPage(result, i, query.PageSize))
}

// GetById 根据id获取任务详情
func (a CronJobApi) GetById(c *gin.Context) {
	jobId := c.Param("jobId")
	info := a.cronJobService.GetInfo(common.StringToInt64(jobId))
	resp.OK(c, util.DeepCopy(info, &response.CronJobResponse{}))
}

// Edit 修改定时任务
func (a CronJobApi) Edit(c *gin.Context) {
	record := entity.SysCronJob{}
	if c.Bind(&record) != nil {
		resp.ParamError(c)
		return
	}
	if a.cronJobService.Edit(record) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Add 新增定时任务
func (a CronJobApi) Add(c *gin.Context) {
	dto := entity.SysCronJob{}
	if c.Bind(&dto) != nil {
		resp.ParamError(c)
		return
	}
	record := a.cronJobService.GetInfoByAlias(dto.FuncAlias)
	if record.JobId > 0 {
		resp.Error(c, "任务已存在!")
	} else {
		a.cronJobService.Save(dto)
		cronjob.AppendCronFunc(dto.JobCron, dto.FuncAlias, dto.Status)
		resp.OK(c)
	}
}

// Remove 删除定时任务
func (a CronJobApi) Remove(c *gin.Context) {
	jobId := common.StringToInt64(c.Param("jobId"))
	funcAlias := c.Param("funcAlias")

	if a.cronJobService.GetInfo(jobId).Status == "1" {
		resp.Error(c, "任务运行中, 无法删除!")
		return
	}

	cronjob.RemoveCronFunc(funcAlias)
	a.cronJobService.Remove([]int64{jobId})
	resp.OK(c)
}

// ChangeStatus 改变任务状态
func (a CronJobApi) ChangeStatus(c *gin.Context) {
	record := entity.SysCronJob{}
	if c.Bind(&record) != nil {
		resp.ParamError(c)
		return
	}
	if a.cronJobService.Edit(record) > 0 {
		if record.Status == "1" {
			cronjob.StartCronFunc(record.FuncAlias)
		} else {
			cronjob.StopCronFunc(record.FuncAlias)
		}
	}
	resp.OK(c)
}
