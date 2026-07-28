package refs

import (
	"context"
	"cutego/pkg/config"
	"cutego/pkg/logger"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// 配置数据库
func InitDatabase() {
	logger.SugaredLogger.Infoln("database init start...")
	var err error
	// 配置mysql数据库
	ds := config.AppEnvConfig.DataSource
	jdbc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		ds.Username,
		ds.Password,
		ds.Host,
		ds.Port,
		ds.Database,
		ds.Charset)
	SqlDB, err = xorm.NewEngine(ds.DbType, jdbc)
	if err != nil {
		logger.SugaredLogger.Fatalf("db error: %#v\n", err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		logger.SugaredLogger.Fatalf("db connect error: %#v\n", err.Error())
	}
	SqlDB.SetMaxIdleConns(ds.MaxIdleSize)
	SqlDB.SetMaxOpenConns(ds.MaxOpenSize)
	ctx, cancel := context.WithCancel(context.Background())
	pingCancel = cancel
	timer := time.NewTicker(time.Minute * 10)
	go func(x *xorm.Engine) {
		for {
			select {
			case <-timer.C:
				err = x.Ping()
				if err != nil {
					logger.SugaredLogger.Fatalf("db connect error: %#v\n", err.Error())
				}
			case <-ctx.Done():
				timer.Stop()
				return
			}
		}
	}(SqlDB)
	SqlDB.ShowSQL(true)
	// 开启缓存
	SqlDB.SetDefaultCacher(xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000))
	// 切换标准时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic("加载时区异常, " + err.Error())
	}
	SqlDB.SetTZLocation(location)
	SqlDB.SetTZDatabase(location)
	logger.SugaredLogger.Infoln("database init end...")
}
