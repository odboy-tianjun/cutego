package v1

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/entity"
	"cutego/modules/core/service"
	"cutego/pkg/page"
	"cutego/pkg/resp"
	"cutego/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type NoticeApi struct {
	noticeService service.NoticeService
}

// List 查询集合
func (a NoticeApi) List(c *gin.Context) {
	query := request.NoticeQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	find, i := a.noticeService.FindPage(query)
	resp.OK(c, page.Page{
		List:  find,
		Total: i,
		Size:  query.PageSize,
	})
}

// Add 添加公告
func (a NoticeApi) Add(c *gin.Context) {
	notice := entity.SysNotice{}
	if c.Bind(&notice) != nil {
		resp.ParamError(c)
		return
	}
	notice.CreateBy = util.GetUserInfo(c).UserName
	if a.noticeService.Save(notice) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Delete 删除
func (a NoticeApi) Delete(c *gin.Context) {
	ids := strings.Split(c.Param("ids"), ",")
	idList := make([]int64, 0)
	for _, s := range ids {
		id, _ := strconv.ParseInt(s, 10, 64)
		idList = append(idList, id)
	}
	if a.noticeService.Remove(idList) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Get 查询
func (a NoticeApi) Get(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	resp.OK(c, a.noticeService.GetById(id))
}

// Edit 修改
func (a NoticeApi) Edit(c *gin.Context) {
	notice := entity.SysNotice{}
	if c.Bind(&notice) != nil {
		resp.ParamError(c)
		return
	}
	notice.UpdateTime = time.Now()
	notice.UpdateBy = util.GetUserInfo(c).UserName
	if a.noticeService.Edit(notice) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}
