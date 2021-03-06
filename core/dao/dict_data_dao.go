package dao

import (
	"cutego/core/api/v1/request"
	"cutego/core/entity"
	"cutego/pkg/common"
	"cutego/pkg/page"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
)

type DictDataDao struct {
}

func (d *DictDataDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_dict_data")
}

// SelectByDictType 根据字典类型查询字典数据
// @Param dictType string 字典类型
// @Return []entity.SysDictData
func (d *DictDataDao) SelectByDictType(dictType string) []entity.SysDictData {
	data := make([]entity.SysDictData, 0)
	session := d.sql(SqlDB.NewSession())
	err := session.Where("status = '0' ").And("dict_type = ?", dictType).OrderBy("dict_sort").Asc("dict_sort").
		Find(&data)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return data
}

// GetDiceDataAll 查询所有字典数据
// @Return *[]entity.SysDictData
func (d DictDataDao) GetDiceDataAll() *[]entity.SysDictData {
	session := d.sql(SqlDB.NewSession())
	data := make([]entity.SysDictData, 0)
	err := session.Where("status = '0' ").OrderBy("dict_sort").Asc("dict_sort").
		Find(&data)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &data
}

// SelectPage 查询集合数据
// @Param query request.DiceDataQuery
// @Return *[]entity.SysDictData
// @Return 总行数
func (d *DictDataDao) SelectPage(query request.DiceDataQuery) (*[]entity.SysDictData, int64) {
	list := make([]entity.SysDictData, 0)
	session := SqlDB.NewSession().Table("sys_dict_data").OrderBy("dict_sort").Asc("dict_sort")
	if gotool.StrUtils.HasNotEmpty(query.DictType) {
		session.And("dict_type = ?", query.DictType)
	}
	if gotool.StrUtils.HasNotEmpty(query.DictLabel) {
		session.And("dict_label like concat('%', ?, '%')", query.DictLabel)
	}
	if gotool.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&list)
	if err != nil {
		common.ErrorLog(err)
		return nil, 0
	}
	return &list, total
}

// SelectByDictCode 根据dictCode查询字典数据
// @Param dictCode int64
// @Return *entity.SysDictData
func (d *DictDataDao) SelectByDictCode(dictCode int64) *entity.SysDictData {
	data := entity.SysDictData{}
	session := SqlDB.NewSession()
	_, err := session.Where("dict_code = ?", dictCode).Get(&data)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &data
}

// Insert 添加字典数据
// @Param data entity.SysDictData
// @Return 新增的行数
func (d *DictDataDao) Insert(data entity.SysDictData) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&data)
	if err != nil {
		session.Rollback()
		common.ErrorLog(err)
		return 0
	}
	session.Commit()
	return insert
}

// Delete 删除字典数据
func (d *DictDataDao) Delete(codes []int64) bool {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.In("dict_code", codes).Delete(&entity.SysDictData{})
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// 修改字典数据
func (d *DictDataDao) Update(data entity.SysDictData) bool {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Where("dict_code = ?", data.DictCode).Update(&data)
	if err != nil {
		session.Rollback()
		common.ErrorLog(err)
		return false
	}
	session.Commit()
	return true
}
