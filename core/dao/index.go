package dao

import (
	models2 "cutego/core/entity"
	"cutego/pkg/common"
	"cutego/pkg/config"
	"cutego/pkg/constant"
	redisTool "cutego/pkg/redispool"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

// X 全局DB
var (
	SqlDB   *xorm.Engine
	RedisDB *redisTool.RedisClient
)

func initDatabase() {
	var err error
	// 配置mysql数据库
	ds := config.AppEnvConfig.DataSource
	jdbc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		ds.Username,
		ds.Password,
		ds.Host,
		ds.Port,
		ds.Database,
		ds.Charset,
	)
	SqlDB, _ = xorm.NewEngine(ds.DbType, jdbc)
	if err != nil {
		common.FatalfLog("db error: %#v\n", err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		common.FatalfLog("db connect error: %#v\n", err.Error())
	}
	SqlDB.SetMaxIdleConns(ds.MaxIdleSize)
	SqlDB.SetMaxOpenConns(ds.MaxOpenSize)
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				common.FatalfLog("db connect error: %#v\n", err.Error())
			}
		}
	}(SqlDB)
	SqlDB.ShowSQL(true)
}
func initRedis() {
	// 配置redis数据库
	RedisDB = redisTool.NewRedis()
}
func init() {
	initDatabase()
	initRedis()
	cacheInitDataToRedis()
}

// 初始化缓存数据
func cacheInitDataToRedis() {
	initDict()
	initConfig()
}

func initDict() {
	// 查询字典类型数据
	dictTypeDao := new(DictTypeDao)
	typeAll := dictTypeDao.SelectAll()
	// 所有字典数据
	d := new(DictDataDao)
	listData := d.GetDiceDataAll()
	for _, dictType := range typeAll {
		dictData := make([]map[string]interface{}, 0)
		for _, data := range *listData {
			if dictType.DictType == data.DictType {
				dictData = append(dictData, map[string]interface{}{
					"dictCode":  data.DictCode,
					"dictSort":  data.DictSort,
					"dictLabel": data.DictLabel,
					"dictValue": data.DictValue,
					"isDefault": data.IsDefault,
					"remark":    data.Remark,
				})
			}
		}
		RedisDB.SET(constant.RedisConst{}.GetRedisDictKey()+dictType.DictType, common.StructToJson(dictData))
	}
}

func initConfig() {
	// 查询配置数据存入到缓存中
	configDao := new(ConfigDao)
	configSession := configDao.sql(SqlDB.NewSession())
	configs := make([]*models2.SysConfig, 0)
	err := configSession.Find(&configs)
	if err != nil {
		common.ErrorLog(err)
		return
	}
	for _, sysConfig := range configs {
		RedisDB.SET(constant.RedisConst{}.GetRedisConfigKey()+sysConfig.ConfigKey, common.StructToJson(map[string]interface{}{
			"configId":    sysConfig.ConfigId,
			"configName":  sysConfig.ConfigName,
			"configKey":   sysConfig.ConfigKey,
			"configValue": sysConfig.ConfigValue,
			"configType":  sysConfig.ConfigType,
			"remark":      sysConfig.Remark,
		}))
	}
}
