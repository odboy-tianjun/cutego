package refs

import (
	"cutego/pkg/config"
	"cutego/pkg/logging"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

// 配置数据库
func init() {
	logging.InfoLog("database init start...")
	var err error
	// 配置mysql数据库
	ds := config.AppEnvConfig.DataSource
	jdbc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		ds.Username,
		ds.Password,
		ds.Host,
		ds.Port,
		ds.Database,
		ds.Charset)
	SqlDB, err = xorm.NewEngine(ds.DbType, jdbc)
	if err != nil {
		logging.FatalfLog("db error: %#v\n", err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		logging.FatalfLog("db connect error: %#v\n", err.Error())
	}
	SqlDB.SetMaxIdleConns(ds.MaxIdleSize)
	SqlDB.SetMaxOpenConns(ds.MaxOpenSize)
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				logging.FatalfLog("db connect error: %#v\n", err.Error())
			}
		}
	}(SqlDB)
	SqlDB.ShowSQL(true)
	// 开启缓存
	SqlDB.SetDefaultCacher(xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000))
	logging.InfoLog("database init end...")
}
