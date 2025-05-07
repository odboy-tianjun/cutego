package page

import (
	"cutego/pkg/logging"
	"github.com/go-xorm/xorm"
)

// Page 分页结构体
type Page struct {
	Size  int         `json:"size"`  // 显示条数
	Total int64       `json:"total"` // 总条数
	List  interface{} `json:"list"`  // 数据
}

type Start struct {
}

// StartSize 获取分页偏移量
func StartSize(pageNum int, size int) int {
	if pageNum == 0 {
		pageNum = 1
	}
	if size == 0 {
		size = 10
	}
	num := (pageNum - 1) * size
	return num
}

// GetTotal 获取总条数
func GetTotal(engine *xorm.Session, args ...interface{}) (int64, error) {
	if args != nil {
		engine.Table(args)
	}
	count, err := engine.Count()
	if err != nil {
		logging.ErrorLog(err.Error())
		return 0, err
	}
	return count, nil
}
