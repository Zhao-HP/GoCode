package user

import (
	"github.com/Zhao-HP/GoCode/web/gin/config"
	"github.com/Zhao-HP/GoCode/web/gin/model/user"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func register(req user.RegisterReq) (*user.TbUser, error) {

	tbUser := &user.TbUser{
		Phone: req.Phone,
	}

	liveRecord := config.LiveRecordConn
	tx := liveRecord.Find(&tbUser)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}
	tbUser.Uid = genUid()
	tbUser.CreateAt = time.Now()
	tbUser.UpdateAt = time.Now()

	config.Log.Info("注册信息: ", tbUser)
	tx = liveRecord.Create(&tbUser)
	if tx.Error != nil {
		config.Log.Error("用户注册失败: ", tx.Error)
		return nil, tx.Error
	}
	return tbUser, nil
}

func login() {

}

func findUserByUid(uid string) {

}
