package v1

import (
	"cutego/modules/core/api/v1/request"
	cache2 "cutego/modules/core/cache"
	"cutego/modules/core/entity"
	service2 "cutego/modules/core/service"
	"cutego/pkg/cache"
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
)

type ConfigApi struct {
	configService service2.ConfigService
	logService    service2.LogService
}

// GetConfigValue 根据参数键名查询参数值
func (a ConfigApi) GetConfigValue(c *gin.Context) {
	param := c.Query("key")
	if !gotool.StrUtils.HasEmpty(param) {
		key := a.configService.GetConfigKey(param)
		c.JSON(http.StatusOK, resp.Success(key.ConfigValue, ""))
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数不合法"))
	}
}

// List 查询设置列表
func (a ConfigApi) List(c *gin.Context) {
	a.logService.LogStart(c)
	a.logService.LogToDB(c, "日志测试 start")

	query := request.ConfigQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	find, i := a.configService.FindPage(query)
	resp.OK(c, page.Page{
		List:  find,
		Total: i,
		Size:  query.PageSize,
	})
	a.logService.LogToDB(c, "日志测试 end")
}

// Add 添加数据
func (a ConfigApi) Add(c *gin.Context) {
	config := entity.SysConfig{}
	if c.Bind(&config) != nil {
		resp.ParamError(c)
		return
	}
	// 检验key是否存在
	if a.configService.CheckConfigKeyUnique(config) {
		resp.Error(c, "新增参数'"+config.ConfigName+"'失败, 参数键名已存在")
	}
	config.CreateBy = util.GetUserInfo(c).UserName
	if a.configService.Save(config) > 0 {
		// 进行缓存数据添加
		cache2.SetRedisConfig(config)
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Get 查询数据
func (a ConfigApi) Get(c *gin.Context) {
	param := c.Param("configId")
	configId, _ := strconv.ParseInt(param, 10, 64)
	resp.OK(c, a.configService.GetInfo(configId))
}

// Edit 修改数据
func (a ConfigApi) Edit(c *gin.Context) {
	config := entity.SysConfig{}
	if c.Bind(&config) != nil {
		resp.ParamError(c)
		return
	}
	// 检验key是否存在
	if a.configService.CheckConfigKeyUnique(config) {
		resp.Error(c, "修改参数'"+config.ConfigName+"'失败, 参数键名已存在")
	}
	config.UpdateBy = util.GetUserInfo(c).UserName
	if a.configService.Edit(config) > 0 {
		cache2.SetRedisConfig(config)
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Delete 删除数据
func (a ConfigApi) Delete(c *gin.Context) {
	ids := strings.Split(c.Param("ids"), ",")
	list := make([]int64, 0)
	for _, id := range ids {
		parseInt, _ := strconv.ParseInt(id, 10, 64)
		list = append(list, parseInt)
	}
	// 进行校验, 查看是否可以删除
	byIds := a.configService.CheckConfigByIds(list)
	for _, config := range *byIds {
		if config.ConfigType == "Y" {
			resp.Error(c, "内置参数"+config.ConfigName+"不能删除")
			return
		}
	}
	// 进行删除
	if a.configService.Remove(list) {
		// 刷新缓存
		strs := make([]string, 0)
		for _, config := range *byIds {
			strs = append(strs, config.ConfigKey)
		}
		cache.RemoveList(strs)
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// RefreshCache 刷新缓存
func (a ConfigApi) RefreshCache(c *gin.Context) {
	all := a.configService.FindAll()
	for _, sysConfig := range *all {
		cache2.RemoveRedisConfig(sysConfig.ConfigKey)
		cache2.SetRedisConfig(sysConfig)
	}
}

// Export 导出数据
func (a ConfigApi) Export(c *gin.Context) {
	query := request.ConfigQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	find, _ := a.configService.FindPage(query)
	items := make([]interface{}, 0)
	for _, config := range *find {
		items = append(items, config)
	}
	_, files := excels.ExportExcel(items, "配置表")
	file.DownloadExcel(c, files)
}
