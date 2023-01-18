package service

import (
	"bytes"
	"cutego/modules/core/api/v1/request"
	"cutego/modules/core/dao"
	"cutego/modules/core/dataobject"
	"github.com/druidcaesa/gotool"
)

type PostService struct {
	postDao dao.PostDao
}

// FindAll 查询所有岗位业务方法
func (s PostService) FindAll() []*dataobject.SysPost {
	return s.postDao.SelectAll()
}

// FindPostListByUserId 根据用户id查询岗位id集合
func (s PostService) FindPostListByUserId(userId int64) *[]int64 {
	return s.postDao.SelectPostListByUserId(userId)
}

// FindList 查询岗位分页列表
func (s PostService) FindPage(query request.PostQuery) (*[]dataobject.SysPost, int64) {
	return s.postDao.SelectPage(query)
}

// CheckPostNameUnique 校验岗位名称是否存在
func (s PostService) CheckPostNameUnique(post dataobject.SysPost) bool {
	return s.postDao.CheckPostNameUnique(post) > 0
}

// CheckPostCodeUnique 校验岗位编码是否存在
func (s PostService) CheckPostCodeUnique(post dataobject.SysPost) bool {
	return s.postDao.CheckPostCodeUnique(post) > 0
}

// Save 添加岗位数据
func (s PostService) Save(post dataobject.SysPost) bool {
	return s.postDao.Insert(post) > 0
}

// GetPostById 根据id查询岗位数据
func (s PostService) GetPostById(id int64) *dataobject.SysPost {
	post := dataobject.SysPost{
		PostId: id,
	}
	return s.postDao.GetPostById(post)
}

// Remove 批量删除岗位信息
func (s PostService) Remove(ids []int64) bool {
	return s.postDao.Delete(ids) > 0
}

// Edit 修改岗位数据
func (s PostService) Edit(post dataobject.SysPost) bool {
	return s.postDao.Update(post)
}

// FindPostByUserName 获取岗位数据
func (s PostService) FindPostByUserName(name string) string {
	list := s.postDao.SelectPostByUserName(name)
	var buffer bytes.Buffer
	var postName string
	for _, post := range *list {
		buffer.WriteString(post.PostName)
		buffer.WriteString(",")
	}
	s2 := buffer.String()
	if gotool.StrUtils.HasNotEmpty(s2) {
		postName = s2[0:(len(s2) - 1)]
	}
	return postName
}
