package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/pkg/common"
	"cutego/pkg/page"
	"cutego/refs"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
)

type CronJobDao struct {
}

func (d CronJobDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_cron_job")
}

// SelectPage 分页查询数据
func (d CronJobDao) SelectPage(query request.CronJobQuery) ([]dataobject.SysCronJob, int64) {
	configs := make([]dataobject.SysCronJob, 0)
	session := d.sql(refs.SqlDB.NewSession())
	if gotool.StrUtils.HasNotEmpty(query.JobName) {
		session.And("job_name like concat('%', ?, '%')", query.JobName)
	}
	if gotool.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&configs)
	if err != nil {
		common.ErrorLog(err)
		return nil, 0
	}
	return configs, total
}

// Insert 添加数据
func (d CronJobDao) Insert(config dataobject.SysCronJob) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&config)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}

// SelectById 查询数据
func (d CronJobDao) SelectById(id int64) *dataobject.SysCronJob {
	config := dataobject.SysCronJob{}
	session := d.sql(refs.SqlDB.NewSession())
	_, err := session.Where("job_id = ?", id).Get(&config)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &config
}

// Update 修改数据
func (d CronJobDao) Update(config dataobject.SysCronJob) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("job_id = ?", config.JobId).Update(&config)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return update
}

// Delete 删除数据
func (d CronJobDao) Delete(list []int64) bool {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.In("job_id", list).Delete(&dataobject.SysCronJob{})
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// SelectByFuncAlias 通过方法别名获取任务详情
func (d CronJobDao) SelectByFuncAlias(funcAlias string) *dataobject.SysCronJob {
	config := dataobject.SysCronJob{}
	session := d.sql(refs.SqlDB.NewSession())
	_, err := session.Where("func_alias = ?", funcAlias).Get(&config)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &config
}

// SelectAll 查找所有启用状态的任务
func (d CronJobDao) SelectAll() ([]dataobject.SysCronJob, int) {
	configs := make([]dataobject.SysCronJob, 0)
	session := d.sql(refs.SqlDB.NewSession())
	err := session.Where("status = ?", 1).Find(&configs)
	if err != nil {
		common.ErrorLog(err)
		return nil, 0
	}
	return configs, len(configs)
}
