package service

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dao"
	"cutego/modules/core/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
)

const CURRENT_UID_KEY = "CURRENT_LOG_UID"
const DEFAULT_UID_KEY = "DEFAULT_LOG_UID"

type LogService struct {
	logDao dao.LogDao
}

// FindPage 分页查询数据
func (s LogService) FindPage(query request.LogQuery) ([]entity.SysLog, int64) {
	return s.logDao.SelectPage(query)
}

// Save 添加数据
func (s LogService) save(config entity.SysLog) int64 {
	return s.logDao.Insert(config)
}

// LogToDB 记录日志
func (s LogService) LogToDB(c *gin.Context, content string) {
	uid, exists := c.Get(CURRENT_UID_KEY)
	if !exists {
		uid = DEFAULT_UID_KEY
	}
	uidStr := uid.(string)
	s.save(entity.SysLog{Uid: uidStr, Content: content})
}

// 开始记录日志前调用(只调用一次)
func (s LogService) LogStart(c *gin.Context) {
	c.Set(CURRENT_UID_KEY, uuid.New())
}
