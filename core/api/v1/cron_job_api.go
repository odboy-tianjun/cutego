package v1

import (
	"cutego/core/api/v1/request"
	"cutego/core/entity"
	"cutego/core/service"
	"cutego/pkg/cache"
	"cutego/pkg/page"
	"cutego/pkg/resp"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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
	resp.OK(c, page.Page{
		List:  find,
		Total: i,
		Size:  query.PageSize,
	})
}

// Edit 修改定时任务
func (a CronJobApi) Edit(c *gin.Context) {
	dictType := entity.SysDictType{}
	if c.Bind(&dictType) != nil {
		resp.ParamError(c)
		return
	}
	//检验定时任务是否存在
	if a.dictTypeService.CheckDictTypeUnique(dictType) {
		resp.Error(c, "修改字典'"+dictType.DictName+"'失败, 定时任务已存在")
		return
	}
	//修改数据
	if a.dictTypeService.Edit(dictType) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Add 新增定时任务
func (a CronJobApi) Add(c *gin.Context) {
	dictType := entity.SysDictType{}
	if c.Bind(&dictType) != nil {
		resp.ParamError(c)
		return
	}
	//检验定时任务是否存在
	if a.dictTypeService.CheckDictTypeUnique(dictType) {
		resp.Error(c, "新增字典'"+dictType.DictName+"'失败, 定时任务已存在")
		return
	}
	//新增定时任务
	if a.dictTypeService.Save(dictType) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Remove 批量删除定时任务
func (a CronJobApi) Remove(c *gin.Context) {
	param := c.Param("dictId")
	split := strings.Split(param, ",")
	ids := make([]int64, 0)
	types := make([]string, 0)
	for _, s := range split {
		parseInt, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, parseInt)
	}
	//校验定时任务是否使用
	for _, id := range ids {
		dictType := a.dictTypeService.GetById(id)
		if len(a.dictDataService.FindByDictType(dictType.DictType)) > 0 {
			resp.Error(c, dictType.DictName+"已分配,不能删除")
			return
		}
		types = append(types, dictType.DictType)
	}
	//批量删除
	if a.dictTypeService.Remove(ids) {
		//从缓存中删除数据
		cache.RemoveList(types)
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// 改变任务状态
func (a CronJobApi) ChangeStatus(context *gin.Context) {

}
