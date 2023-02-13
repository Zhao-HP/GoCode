package router

import (
	"github.com/Zhao-HP/GoCode/web/gin/module/user"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	user.InitRouter(router)
}
