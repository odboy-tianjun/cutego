package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/pkg/common"
	"cutego/refs"
	"github.com/druidcaesa/gotool"
)

type DeptDao struct {
}

// SelectTree 根据条件查询部门集合
func (d DeptDao) SelectTree(query request.DeptQuery) *[]dataobject.SysDept {
	depts := make([]dataobject.SysDept, 0)
	session := refs.SqlDB.NewSession().Where("del_flag = '0'")
	if query.ParentId > 0 {
		session.And("parent_id = ?", query.ParentId)
	}
	if !gotool.StrUtils.HasEmpty(query.DeptName) {
		session.And("dept_name like concat('%', ?, '%')", query.DeptName)
	}
	if !gotool.StrUtils.HasEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	err := session.OrderBy("parent_id").OrderBy("order_num").Find(&depts)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &depts
}

// SelectDeptListByRoleId 根据角色ID查询部门树信息
func (d DeptDao) SelectDeptListByRoleId(id int64, strictly bool) *[]int64 {
	list := make([]int64, 0)
	session := refs.SqlDB.NewSession().Table([]string{"sys_dept", "d"}).Cols("d.dept_id")
	session.Join("LEFT", []string{"sys_role_dept", "rd"}, "d.dept_id = rd.dept_id").
		Where("rd.role_id = ?", id)
	if strictly {
		session.And("d.dept_id not in (select d.parent_id from sys_dept d inner join sys_role_dept rd on d.dept_id = rd.dept_id and rd.role_id = ?)", id)
	}
	err := session.OrderBy("d.parent_id").OrderBy("d.order_num").Find(&list)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &list
}

// GetList 查询部门列表
func (d DeptDao) GetList(query request.DeptQuery) *[]dataobject.SysDept {
	list := make([]dataobject.SysDept, 0)
	session := refs.SqlDB.NewSession().OrderBy("parent_id").OrderBy("order_num")
	session.Where("del_flag = '0'")
	if query.ParentId > 0 {
		session.And("parent_id = ?", query.ParentId)
	}
	if gotool.StrUtils.HasNotEmpty(query.DeptName) {
		session.And("dept_name like concat('%', ?, '%')", query.DeptName)
	}
	if gotool.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	err := session.Find(&list)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &list
}

// SelectDeptById 根据部门编号获取详细信息
func (d DeptDao) SelectDeptById(id int) *dataobject.SysDept {
	dept := dataobject.SysDept{}
	_, err := refs.SqlDB.NewSession().Where("dept_id = ?", id).Get(&dept)
	if err != nil {
		common.ErrorLog(err)
		return nil
	}
	return &dept
}

// Insert 添加部门数据
func (d DeptDao) Insert(dept dataobject.SysDept) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&dept)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}

// CheckDeptNameUnique 校验部门名称是否唯一
func (d DeptDao) CheckDeptNameUnique(dept dataobject.SysDept) int64 {
	session := refs.SqlDB.NewSession()
	count, err := session.Table("sys_dept").Cols("dept_id").Where("dept_name=?", dept.DeptName).And("parent_id = ?", dept.ParentId).Limit(1).Count()
	if err != nil {
		common.ErrorLog(err)
		return 1
	}
	return count

}

// HasChildByDeptId 是否存在部门子节点
func (d DeptDao) HasChildByDeptId(id int) int64 {
	count, _ := refs.SqlDB.NewSession().Table("sys_dept").Cols("dept_id").Where("parent_id = ?", id).
		And("del_flag = '0'").Limit(1).Count()
	return count
}

// CheckDeptExistUser 查询部门是否存在用户
func (d DeptDao) CheckDeptExistUser(id int) int64 {
	count, _ := refs.SqlDB.NewSession().Table("sys_user").Cols("user_id").Where("dept_id = ?", id).
		And("del_flag = '0'").Count()
	return count
}

// Delete 删除部门
func (d DeptDao) Delete(id int) int64 {
	dept := dataobject.SysDept{
		DeptId: id,
	}
	session := refs.SqlDB.NewSession()
	session.Begin()
	i, err := session.Where("dept_id = ?", id).Delete(&dept)
	if err != nil {
		session.Rollback()
		common.ErrorLog(err)
		return 0
	}
	session.Commit()
	return i
}

// Update 更新部门
func (d DeptDao) Update(dept dataobject.SysDept) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("dept_id = ?", dept.DeptId).Update(&dept)
	if err != nil {
		session.Rollback()
		common.ErrorLog(err)
		return 0
	}
	session.Commit()
	return update
}
