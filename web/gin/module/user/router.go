package user

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {

	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", Login)
		userGroup.POST("register", Register)
		userGroup.GET("/findUserByUid", FindUserByUid)
	}

}
