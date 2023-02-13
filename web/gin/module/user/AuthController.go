package user

import (
	util "github.com/Zhao-HP/GoCode/utils"
	"github.com/Zhao-HP/GoCode/web/gin/config"
	"github.com/Zhao-HP/GoCode/web/gin/model/constant"
	"github.com/Zhao-HP/GoCode/web/gin/model/result"
	"github.com/Zhao-HP/GoCode/web/gin/model/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(ctx *gin.Context) {
	var req user.RegisterReq
	if err := ctx.ShouldBindJSON(req); err != nil {
		result.Error(ctx, constant.DefaultErrorCode, constant.ParamError)
		return
	}

}

func Register(ctx *gin.Context) {
	var req user.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		result.Error(ctx, constant.DefaultErrorCode, constant.ParamError)
		return
	}
	if util.StringIsEmpty(req.Phone) {
		result.Error(ctx, constant.DefaultErrorCode, "手机号不能为空")
		return
	}

	tbUser, err := register(req)
	if err != nil {
		result.Error(ctx, constant.DefaultErrorCode, err.Error())
		return
	}
	result.Success(ctx, tbUser)
}

func FindUserByUid(ctx *gin.Context) {
	uid := ctx.Query("uid")
	if len(uid) <= 0 {
		result.Error(ctx, constant.DefaultErrorCode, constant.ParamError)
		return
	}

	conn := config.LiveRecordConn
	tbUser := user.TbUser{}

	conn.Where("uid = ?", uid).Find(&tbUser)
	result.Success(ctx, tbUser)
}

func genUid() string {
	// 生成纯数字的UID
	uid, _ := uuid.NewUUID()
	return uid.String()
}
