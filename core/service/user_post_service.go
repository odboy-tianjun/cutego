package service

import (
	"cutego/core/dao"
)

type UserPostService struct {
	userPostDao dao.UserPostDao
}

// CountUserPostById 统计岗位数据数量
func (s UserPostService) CountUserPostById(ids []int64) int64 {
	for _, id := range ids {
		if s.userPostDao.CountById(id) > 0 {
			return id
		}
	}
	return 0
}
