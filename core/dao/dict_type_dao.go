package dao

import (
	"cutego/core/api/v1/request"
	"cutego/core/entity"
	"cutego/pkg/common"
	"cutego/pkg/page"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
)

type DictTypeDao struct {
}

func (d DictTypeDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_dict_type")
}

// SelectAll 查询所有字典类型数据
func (d DictTypeDao) SelectAll() []*entity.SysDictType {
	types := make([]*entity.SysDictType, 0)
	err := d.sql(SqlDB.NewSession()).Where("status = '0'").Find(&types)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return types
}

// SelectPage 分页查询字典类型数据
func (d DictTypeDao) SelectPage(query request.DictTypeQuery) (*[]entity.SysDictType, int64) {
	list := make([]entity.SysDictType, 0)
	session := SqlDB.NewSession().Table("sys_dict_type")
	if gotool.StrUtils.HasNotEmpty(query.DictName) {
		session.And("dict_name like concat('%', ?, '%')", query.DictName)
	}
	if gotool.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	if gotool.StrUtils.HasNotEmpty(query.DictType) {
		session.And("AND dict_type like concat('%', ?, '%')", query.DictType)
	}
	if gotool.StrUtils.HasNotEmpty(query.BeginTime) {
		session.And("date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d')", query.BeginTime)
	}
	if gotool.StrUtils.HasNotEmpty(query.EndTime) {
		session.And("date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d')", query.EndTime)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&list)
	if err != nil {
		common.ErrorLog(err)
		return nil, 0
	}
	return &list, total
}

// SelectById 根据id查询字典类型数据
func (d DictTypeDao) SelectById(id int64) *entity.SysDictType {
	dictType := entity.SysDictType{}
	_, err := SqlDB.NewSession().Where("dict_id = ?", id).Get(&dictType)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &dictType
}

// CheckDictTypeUnique 检验字典类型是否存在
func (d DictTypeDao) CheckDictTypeUnique(dictType entity.SysDictType) int64 {
	session := SqlDB.NewSession().Table("sys_dict_type")
	if dictType.DictId > 0 {
		session.And("dict_id != ?", dictType.DictId)
	}
	count, err := session.Where("dict_type = ?", dictType.DictType).Cols("dict_id").Count()
	if err != nil {
		common.ErrorLog(err)
		return 0
	}
	return count
}

// Update 修改字典
func (d DictTypeDao) Update(dictType entity.SysDictType) bool {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Where("dict_id = ?", dictType.DictId).Update(&dictType)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// Insert 新增字典类型
func (d DictTypeDao) Insert(dictType entity.SysDictType) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&dictType)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}

// Delete 批量删除
func (d DictTypeDao) Delete(ids []int64) bool {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.In("dict_id", ids).Delete(entity.SysDictType{})
	if err != nil {
		session.Rollback()
		common.ErrorLog(err)
		return false
	}
	session.Commit()
	return true
}
