package dao

import (
	"cutego/modules/core/entity"
	"cutego/pkg/common"
)

type UserPostDao struct {
}

// BatchInsert 批量新增用户岗位信息
func (d UserPostDao) BatchInsert(posts []entity.SysUserPost) {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Table(entity.SysUserPost{}.TableName()).Insert(&posts)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// Delete 删除用户和岗位关系
func (d UserPostDao) Delete(id int64) {
	post := entity.SysUserPost{
		UserId: id,
	}
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&post)
	if err != nil {
		common.ErrorLog(err)
		session.Rollback()
	}
	session.Commit()
}

// CountById 通过岗位ID查询岗位使用数量
func (d UserPostDao) CountById(id int64) int64 {
	count, err := SqlDB.NewSession().Table("sys_user_post").Cols("post_id").Where("post_id = ?", id).Count()
	if err != nil {
		common.ErrorLog(err)
		return 0
	}
	return count
}
