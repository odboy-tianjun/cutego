package dao

import (
	"cutego/core/api/v1/request"
	"cutego/core/entity"
	"cutego/pkg/common"
	"cutego/pkg/page"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
)

type CronJobDao struct {
}

func (d CronJobDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_cron_job")
}

// SelectPage 分页查询数据
func (d CronJobDao) SelectPage(query request.CronJobQuery) ([]entity.SysCronJob, int64) {
	configs := make([]entity.SysCronJob, 0)
	session := d.sql(SqlDB.NewSession())
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
func (d CronJobDao) Insert(config entity.SysCronJob) int64 {
	session := SqlDB.NewSession()
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
func (d CronJobDao) SelectById(id int64) *entity.SysCronJob {
	config := entity.SysCronJob{}
	session := d.sql(SqlDB.NewSession())
	_, err := session.Where("job_id = ?", id).Get(&config)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &config
}

// Update 修改数据
func (d CronJobDao) Update(config entity.SysCronJob) int64 {
	session := SqlDB.NewSession()
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

// Remove 删除数据
func (d CronJobDao) Delete(list []int64) bool {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.In("job_id", list).Delete(&entity.SysCronJob{})
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// 通过方法别名获取任务详情
func (d CronJobDao) SelectByFuncAlias(funcAlias string) *entity.SysCronJob {
	config := entity.SysCronJob{}
	session := d.sql(SqlDB.NewSession())
	_, err := session.Where("func_alias = ?", funcAlias).Get(&config)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &config
}
