package v1

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/api/v1/response"
	"cutego/modules/core/entity"
	"cutego/modules/core/service"
	"cutego/pkg/excels"
	"cutego/pkg/file"
	"cutego/pkg/page"
	"cutego/pkg/resp"
	"cutego/pkg/util"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type DictDataApi struct {
	dictDataService service.DictDataService
}

// GetByType 根据字典类型查询字典数据d
func (a DictDataApi) GetByType(c *gin.Context) {
	param := c.Param("dictType")
	if !gotool.StrUtils.HasEmpty(param) {
		list := a.dictDataService.FindByDictType(param)
		var result []response.DictDataResponse
		for _, data := range list {
			result = append(result, response.DictDataResponse{
				DictCode:  data.DictCode,
				DictLabel: data.DictLabel,
				DictSort:  data.DictSort,
				DictValue: data.DictValue,
				DictType:  data.DictType,
				IsDefault: data.IsDefault,
			})
		}
		c.JSON(http.StatusOK, resp.Success(result))
	}
}

// List 查询字典数据集合
func (a DictDataApi) List(c *gin.Context) {
	query := request.DiceDataQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	list, i := a.dictDataService.FindPage(query)
	resp.OK(c, page.Page{
		List:  list,
		Total: i,
		Size:  query.PageSize,
	})
}

// Get 根据id查询字典数据
func (a DictDataApi) Get(c *gin.Context) {
	param := c.Param("dictCode")
	dictCode, _ := strconv.ParseInt(param, 10, 64)
	resp.OK(c, a.dictDataService.GetByCode(dictCode))
}

// Add 添加字典数据
func (a DictDataApi) Add(c *gin.Context) {
	data := entity.SysDictData{}
	if c.Bind(&data) != nil {
		resp.ParamError(c)
		return
	}
	data.CreateBy = util.GetUserInfo(c).UserName
	if a.dictDataService.Save(data) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Edit 编辑字典数据
func (a DictDataApi) Edit(c *gin.Context) {
	data := entity.SysDictData{}
	if c.Bind(&data) != nil {
		resp.ParamError(c)
		return
	}
	data.UpdateBy = util.GetUserInfo(c).UserName
	data.UpdateTime = time.Now()
	if a.dictDataService.Edit(data) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Delete 删除数据
func (a DictDataApi) Delete(c *gin.Context) {
	param := c.Param("dictCode")
	split := strings.Split(param, ",")
	dictCodeList := make([]int64, 0)
	for _, s := range split {
		diceCode, _ := strconv.ParseInt(s, 10, 64)
		dictCodeList = append(dictCodeList, diceCode)
	}
	if a.dictDataService.Remove(dictCodeList) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Export 导出excel
func (a DictDataApi) Export(c *gin.Context) {
	query := request.DiceDataQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	items := make([]interface{}, 0)
	list, _ := a.dictDataService.FindPage(query)
	for _, data := range *list {
		items = append(items, data)
	}
	_, files := excels.ExportExcel(items, "字典数据表")
	file.DownloadExcel(c, files)
}
