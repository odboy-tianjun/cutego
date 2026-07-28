package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/pkg/logger"
	"cutego/pkg/page"
	"cutego/refs"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
)

type JobLogDao struct {
}

func (d JobLogDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_job_log")
}

// SelectPage 分页查询
func (d JobLogDao) SelectPage(query request.JobLogQuery) ([]dataobject.SysJobLog, int64) {
	logs := make([]dataobject.SysJobLog, 0)
	session := d.sql(refs.SqlDB.NewSession())
	if gotool.StrUtils.HasNotEmpty(query.JobName) {
		session.And("job_name like concat('%', ?, '%')", query.JobName)
	}
	if gotool.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	if gotool.StrUtils.HasNotEmpty(query.BeginTime) {
		session.And("date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d')", query.BeginTime)
	}
	if gotool.StrUtils.HasNotEmpty(query.EndTime) {
		session.And("date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d')", query.EndTime)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.OrderBy("create_time desc").Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&logs)
	if err != nil {
		logger.SugaredLogger.Errorln(err)
		return nil, 0
	}
	return logs, total
}

// DeleteByIds 批量删除
func (d JobLogDao) DeleteByIds(ids []int64) bool {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.In("job_log_id", ids).Delete(&dataobject.SysJobLog{})
	if err != nil {
		logger.SugaredLogger.Errorln(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// Clean 清空所有日志
func (d JobLogDao) Clean() bool {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.Exec("delete from sys_job_log")
	if err != nil {
		logger.SugaredLogger.Errorln(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// Insert 插入日志
func (d JobLogDao) Insert(log dataobject.SysJobLog) {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.Insert(&log)
	if err != nil {
		logger.SugaredLogger.Errorln(err)
		session.Rollback()
		return
	}
	session.Commit()
}