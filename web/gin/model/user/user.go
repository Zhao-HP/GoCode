package user

import "time"

type TbUser struct {
	Id       int       `json:"id"`
	Uid      string    `json:"uid"`
	Phone    string    `json:"phone"`
	NickName string    `json:"nickName"`
	OpenId   string    `json:"openId"`
	UnionId  string    `json:"unionId"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func (TbUser) TableName() string {
	return "tb_users"
}
