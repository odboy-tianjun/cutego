package v1

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/service"
	"cutego/pkg/resp"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type JobLogApi struct {
	jobLogService service.JobLogService
}

// List 查询调度日志列表
func (a JobLogApi) List(c *gin.Context) {
	query := request.JobLogQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	list, total := a.jobLogService.FindPage(query)
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "查询成功",
		"rows":   list,
		"total":  total,
	})
}

// Remove 删除调度日志
func (a JobLogApi) Remove(c *gin.Context) {
	param := c.Param("jobLogId")
	ids := strings.Split(param, ",")
	idList := make([]int64, 0, len(ids))
	for _, s := range ids {
		id, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
		if err == nil {
			idList = append(idList, id)
		}
	}
	if len(idList) > 0 && a.jobLogService.Delete(idList) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Clean 清空调度日志
func (a JobLogApi) Clean(c *gin.Context) {
	if a.jobLogService.Clean() {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}