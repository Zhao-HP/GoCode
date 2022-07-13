package test

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
	"time"
)

type User struct {
	UID        uint      `json:"uid,select(article)"`
	Avatar     string    `json:"avatar,select(article)"`
	Nickname   string    `json:"nickname,select(article|profile)"`
	Sex        int       `json:"sex"`
	VipEndTime time.Time `json:"vip_end_time,select(profile)"`
	Price      string    `json:"price,select(profile)"`
}

func NewUser() User {
	return User{
		UID:        0,
		Avatar:     "https://gimg2.baidu.com",
		Nickname:   "昵称",
		Sex:        1,
		VipEndTime: time.Now().Add(time.Hour * 24 * 365),
		Price:      "19999.9",
	}
}

func TestSelector(t *testing.T) {
	articleUser := filter.SelectMarshal("article", NewUser())
	fmt.Println(articleUser.MustJSON())

	profileUser := filter.SelectMarshal("profile", NewUser())
	fmt.Println(profileUser.MustJSON())

}
