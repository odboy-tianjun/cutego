package dao

import (
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dataobject"
	"cutego/pkg/logging"
	"cutego/pkg/page"
	"cutego/refs"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
)

type PostDao struct {
}

func (d PostDao) sqlSelectJoin(session *xorm.Session) *xorm.Session {
	return session.Table([]string{"sys_post", "p"}).
		Join("LEFT", []string{"sys_user_post", "up"}, "up.post_id = p.post_id").
		Join("LEFT", []string{"sys_user", "u"}, "u.user_id = up.user_id")
}

// SelectAll 查询所有岗位数据, 数据库操作
func (d PostDao) SelectAll() []*dataobject.SysPost {
	session := refs.SqlDB.NewSession()
	posts := make([]*dataobject.SysPost, 0)
	err := session.Find(&posts)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return posts
}

// SelectPostListByUserId 根据用户id查询岗位id集合
func (d PostDao) SelectPostListByUserId(userId int64) *[]int64 {
	var ids []int64
	selectSql := d.sqlSelectJoin(refs.SqlDB.NewSession())
	err := selectSql.Where("u.user_id = ?", userId).Cols("p.post_id").Find(&ids)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &ids
}

// SelectPage 查询岗位分页数据
func (d PostDao) SelectPage(query request.PostQuery) (*[]dataobject.SysPost, int64) {
	posts := make([]dataobject.SysPost, 0)
	session := refs.SqlDB.NewSession().Table(dataobject.SysPost{}.TableName())
	if gotool.StrUtils.HasNotEmpty(query.PostCode) {
		session.And("post_code like concat('%', ?, '%')", query.PostCode)
	}
	if gotool.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	if gotool.StrUtils.HasNotEmpty(query.PostName) {
		session.And("post_name like concat('%', ?, '%')", query.PostName)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&posts)
	if err != nil {
		logging.ErrorLog(err)
		return nil, 0
	}
	return &posts, total
}

// CheckPostNameUnique 校验岗位名称是否存在
func (d PostDao) CheckPostNameUnique(post dataobject.SysPost) int64 {
	session := refs.SqlDB.NewSession().Table("sys_post").Cols("post_id").
		Where("post_name = ?", post.PostName)
	if post.PostId > 0 {
		session.And("post_id != ?", post.PostId)
	}
	count, _ := session.Count()
	return count
}

// CheckPostCodeUnique 校验岗位编码是否存在
func (d PostDao) CheckPostCodeUnique(post dataobject.SysPost) int64 {
	session := refs.SqlDB.NewSession().Table("sys_post").Cols("post_id").
		Where("post_code = ?", post.PostCode)
	if post.PostId > 0 {
		session.And("post_id != ?", post.PostId)
	}
	count, _ := session.Count()
	return count
}

// Insert 添加岗位数据
func (d PostDao) Insert(post dataobject.SysPost) int64 {
	session := refs.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&post)
	if err != nil {
		session.Rollback()
		logging.ErrorLog(err)
		return 0
	}
	session.Commit()
	return insert
}

// GetPostById 根据id查询岗位数据
func (d PostDao) GetPostById(post dataobject.SysPost) *dataobject.SysPost {
	_, err := refs.SqlDB.NewSession().Where("post_id = ?", post.PostId).Get(&post)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &post
}

// Delete 批量删除岗位
func (d PostDao) Delete(posts []int64) int64 {

	session := refs.SqlDB.NewSession()
	session.Begin()
	i, err := session.In("post_id", posts).Delete(&dataobject.SysPost{})
	if err != nil {
		session.Rollback()
		logging.ErrorLog(err)
		return 0
	}
	session.Commit()
	return i
}

// Update 修改岗位数据
func (d PostDao) Update(post dataobject.SysPost) bool {
	session := refs.SqlDB.NewSession()
	session.Begin()
	_, err := session.Where("post_id = ?", post.PostId).Update(&post)
	if err != nil {
		session.Rollback()
		logging.ErrorLog(err)
		return false
	}
	session.Commit()
	return true
}

func (d PostDao) SelectPostByUserName(name string) *[]dataobject.SysPost {
	posts := make([]dataobject.SysPost, 0)
	session := refs.SqlDB.NewSession().Table([]string{dataobject.SysPost{}.TableName(), "p"})
	err := session.Cols("p.post_id", "p.post_name", "p.post_code").
		Join("LEFT", []string{"sys_user_post", "up"}, "up.post_id = p.post_id").
		Join("LEFT", []string{"sys_user", "u"}, "u.user_id = up.user_id").Where("u.user_name = ?", name).Find(&posts)
	if err != nil {
		logging.ErrorLog(err)
		return nil
	}
	return &posts
}
