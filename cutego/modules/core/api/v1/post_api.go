package v1

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/modules/core/service"
	"cutego/pkg/excels"
	"cutego/pkg/file"
	"cutego/pkg/page"
	"cutego/pkg/resp"
	"cutego/pkg/util"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type PostApi struct {
	postService     service.PostService
	userPostService service.UserPostService
}

// List 查询刚问分页数据
func (a PostApi) List(c *gin.Context) {
	query := request.PostQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	list, i := a.postService.FindPage(query)
	resp.OK(c, page.Page{
		List:  list,
		Total: i,
		Size:  query.PageSize,
	})
}

// Add 新增岗位
func (a PostApi) Add(c *gin.Context) {
	post := dataobject.SysPost{}
	if c.Bind(&post) != nil {
		resp.ParamError(c)
		return
	}
	// 校验岗位名称是否存在
	if a.postService.CheckPostNameUnique(post) {
		resp.Error(c, "新增岗位'"+post.PostName+"'失败, 岗位名称已存在")
		return
	}
	// 检验岗位编码是否存在
	if a.postService.CheckPostCodeUnique(post) {
		resp.Error(c, "新增岗位'"+post.PostCode+"'失败, 岗位编码已存在")
		return
	}
	post.CreateBy = util.GetUserInfo(c).UserName
	if a.postService.Save(post) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Get 根据岗位编号获取详细信息
func (a PostApi) Get(c *gin.Context) {
	param := c.Param("postId")
	postId, _ := strconv.ParseInt(param, 10, 64)
	resp.OK(c, a.postService.GetPostById(postId))
}

// Delete 删除岗位数据
func (a PostApi) Delete(c *gin.Context) {
	// 获取postId
	param := c.Param("postId")
	list := make([]int64, 0)
	if gotool.StrUtils.HasNotEmpty(param) {
		strs := strings.Split(param, ",")
		for _, str := range strs {
			postId, _ := strconv.ParseInt(str, 10, 64)
			list = append(list, postId)
		}
	}
	// 判断是否可以删除
	postId := a.userPostService.CountUserPostById(list)
	if postId > 0 {
		post := a.postService.GetPostById(postId)
		resp.Error(c, post.PostName+"岗位已分配, 不允许删除")
		return
	}
	if a.postService.Remove(list) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Edit 修改岗位数据接口
func (a PostApi) Edit(c *gin.Context) {
	post := dataobject.SysPost{}
	if c.Bind(&post) != nil {
		resp.ParamError(c)
		return
	}
	// 校验岗位名称是否存在
	if a.postService.CheckPostNameUnique(post) {
		resp.Error(c, "修改岗位'"+post.PostName+"'失败, 岗位名称已存在")
		return
	}
	// 检验岗位编码是否存在
	if a.postService.CheckPostCodeUnique(post) {
		resp.Error(c, "修改岗位'"+post.PostCode+"'失败, 岗位编码已存在")
		return
	}
	post.UpdateBy = util.GetUserInfo(c).UserName
	post.UpdateTime = time.Now()
	if a.postService.Edit(post) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Export 导出excel
func (a PostApi) Export(c *gin.Context) {
	query := request.PostQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	list, _ := a.postService.FindPage(query)
	excelList := make([]interface{}, 0)
	for _, post := range *list {
		excelList = append(excelList, post)
	}
	_, files := excels.ExportExcel(excelList, "岗位数据表")
	file.DownloadExcel(c, files)
}
