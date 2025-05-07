package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/pkg/cache"
	"cutego/pkg/constant"
	"cutego/pkg/logging"
	"cutego/pkg/page"
	"cutego/refs"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
)

type ConfigDao struct {
}

func (d ConfigDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_config")
}

// SelectByConfigKey 根据键名查询参数配置信息
func (d ConfigDao) SelectByConfigKey(configKey string) *dataobject.SysConfig {
	config := dataobject.SysConfig{}
	_, err := d.sql(refs.SqlDB.NewSession()).Where("config_key = ?", configKey).Get(&config)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &config
}

// SelectPage 分页查询数据
func (d ConfigDao) SelectPage(query request.ConfigQuery) (*[]dataobject.SysConfig, int64) {
	configs := make([]dataobject.SysConfig, 0)
	session := d.sql(refs.SqlDB.NewSession())
	if gotool.StrUtils.HasNotEmpty(query.ConfigName) {
		session.And("config_name like concat('%', ?, '%')", query.ConfigName)
	}
	if gotool.StrUtils.HasNotEmpty(query.ConfigType) {
		session.And("config_type = ?", query.ConfigType)
	}
	if gotool.StrUtils.HasNotEmpty(query.ConfigKey) {
		session.And("config_key like concat('%', ?, '%')", query.ConfigKey)
	}
	if gotool.StrUtils.HasNotEmpty(query.BeginTime) {
		session.And("date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d')", query.BeginTime)
	}
	if gotool.StrUtils.HasNotEmpty(query.EndTime) {
		session.And("date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d')", query.EndTime)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&configs)
	if err != nil {
		logging.ErrorLog(err)
		return nil, 0
	}
	return &configs, total
}

// CheckConfigKeyUnique 校验是否存在
func (d ConfigDao) CheckConfigKeyUnique(config dataobject.SysConfig) int64 {
	session := d.sql(refs.SqlDB.NewSession())
	if config.ConfigId > 0 {
		session.Where("config_id != ?", config.ConfigId)
	}
	count, err := session.And("config_key = ?", config.ConfigKey).Cols("config_id").Count()
	if err != nil {
		logging.ErrorLog(err)
		return 0
	}
	return count
}

// Insert 添加数据
func (d ConfigDao) Insert(config dataobject.SysConfig) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&config)
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}

// SelectById 查询数据
func (d ConfigDao) SelectById(id int64) *dataobject.SysConfig {
	config := dataobject.SysConfig{}
	session := d.sql(refs.SqlDB.NewSession())
	_, err := session.Where("config_id = ?", id).Get(&config)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &config
}

// Update 修改数据
func (d ConfigDao) Update(config dataobject.SysConfig) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("config_id = ?", config.ConfigId).Update(&config)
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return update
}

// CheckConfigByIds 根据id集合查询
func (d ConfigDao) CheckConfigByIds(list []int64) *[]dataobject.SysConfig {
	configs := make([]dataobject.SysConfig, 0)
	err := d.sql(refs.SqlDB.NewSession()).In("config_id", list).Find(&configs)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &configs
}

// Remove 删除数据
func (d ConfigDao) Delete(list []int64) bool {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.In("config_id", list).Delete(&dataobject.SysConfig{})
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// SelectAll 查询所有数据
func (d ConfigDao) SelectAll() *[]dataobject.SysConfig {
	configs := make([]dataobject.SysConfig, 0)
	session := refs.SqlDB.NewSession()
	err := session.Find(&configs)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &configs
}

func init() {
	// 查询配置数据存入到缓存中
	configDao := new(ConfigDao)
	configSession := configDao.sql(refs.SqlDB.NewSession())
	configs := make([]*dataobject.SysConfig, 0)
	err := configSession.Find(&configs)
	if err != nil {
		logging.ErrorLog(err)
		return
	}
	for _, sysConfig := range configs {
		refs.RedisDB.SET(constant.RedisConst{}.GetRedisConfigKey()+sysConfig.ConfigKey, cache.StructToJson(map[string]interface{}{
			"configId":    sysConfig.ConfigId,
			"configName":  sysConfig.ConfigName,
			"configKey":   sysConfig.ConfigKey,
			"configValue": sysConfig.ConfigValue,
			"configType":  sysConfig.ConfigType,
			"remark":      sysConfig.Remark,
		}))
	}
}
