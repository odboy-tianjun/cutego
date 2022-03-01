package entity

import "time"

type SysNotice struct {
	NoticeId      int64     `xorm:"pk autoincr" json:"noticeId"`     // 公告id
	NoticeTitle   string    `xorm:"notice_title" json:"noticeTitle"` // 公告标题
	NoticeType    string    `xorm:"notice_type" json:"noticeType"`   // 公告类型（1通知 2公告）
	NoticeContent string    `json:"noticeContent"`                   // 公告内容
	Status        string    `json:"status"`                          // 公告状态（0正常 1关闭）
	CreateBy      string    `xorm:"varchar(64)" json:"createBy"`     // 创建人
	CreateTime    time.Time `xorm:"created" json:"createTime"`       // 创建时间
	UpdateBy      string    `xorm:"varchar(64)" json:"updateBy"`     // 更新人
	UpdateTime    time.Time `json:"updateTime"`                      // 更新时间
	Remark        string    `xorm:"varchar(500)" json:"remark"`      // 备注
}

func (SysNotice) TableName() string {
	return "sys_notice"
}
