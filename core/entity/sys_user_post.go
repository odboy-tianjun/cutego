package entity

type SysUserPost struct {
	UserId int64 `json:"userId"` // 用户id
	PostId int64 `json:"postId"` // 部门id
}

func (SysUserPost) TableName() string {
	return "sys_user_post"
}
