package constant

const (
	mysqlErrorMsg = "调用MySQL发生异常, %s"
)

type MysqlConst struct{}

// GetMysqlError Mysql异常拼接常量
func (c MysqlConst) GetMysqlError() string {
	return mysqlErrorMsg
}
