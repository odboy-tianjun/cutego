package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/pkg/common"
	"cutego/pkg/constant"
	"cutego/pkg/logging"
	"cutego/pkg/page"
	"cutego/refs"
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
// @Return []dataobject.SysDictData
func (d *DictDataDao) SelectByDictType(dictType string) []dataobject.SysDictData {
	data := make([]dataobject.SysDictData, 0)
	session := d.sql(refs.SqlDB.NewSession())
	err := session.Where("status = '0' ").And("dict_type = ?", dictType).OrderBy("dict_sort").Asc("dict_sort").
		Find(&data)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return data
}

// GetDiceDataAll 查询所有字典数据
// @Return *[]dataobject.SysDictData
func (d DictDataDao) GetDiceDataAll() *[]dataobject.SysDictData {
	session := d.sql(refs.SqlDB.NewSession())
	data := make([]dataobject.SysDictData, 0)
	err := session.Where("status = '0' ").OrderBy("dict_sort").Asc("dict_sort").
		Find(&data)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &data
}

// SelectPage 查询集合数据
// @Param query request.DiceDataQuery
// @Return *[]dataobject.SysDictData
// @Return 总行数
func (d *DictDataDao) SelectPage(query request.DiceDataQuery) (*[]dataobject.SysDictData, int64) {
	list := make([]dataobject.SysDictData, 0)
	session := refs.SqlDB.NewSession().Table("sys_dict_data").OrderBy("dict_sort").Asc("dict_sort")
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
		logging.ErrorLog(err)
		return nil, 0
	}
	return &list, total
}

// SelectByDictCode 根据dictCode查询字典数据
// @Param dictCode int64
// @Return *dataobject.SysDictData
func (d *DictDataDao) SelectByDictCode(dictCode int64) *dataobject.SysDictData {
	data := dataobject.SysDictData{}
	session := refs.SqlDB.NewSession()
	_, err := session.Where("dict_code = ?", dictCode).Get(&data)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &data
}

// Insert 添加字典数据
// @Param data dataobject.SysDictData
// @Return 新增的行数
func (d *DictDataDao) Insert(data dataobject.SysDictData) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&data)
	if err != nil {
		session.Rollback()
		logging.ErrorLog(err)
		return 0
	}
	session.Commit()
	return insert
}

// Delete 删除字典数据
func (d *DictDataDao) Delete(codes []int64) bool {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.In("dict_code", codes).Delete(&dataobject.SysDictData{})
	if err != nil {
		logging.ErrorLog(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// 修改字典数据
func (d *DictDataDao) Update(data dataobject.SysDictData) bool {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.Where("dict_code = ?", data.DictCode).Update(&data)
	if err != nil {
		session.Rollback()
		logging.ErrorLog(err)
		return false
	}
	session.Commit()
	return true
}

func init() {
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
		refs.RedisDB.SET(constant.RedisConst{}.GetRedisDictKey()+dictType.DictType, common.StructToJson(dictData))
	}
}
