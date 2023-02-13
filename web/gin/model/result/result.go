package result

import (
	"github.com/Zhao-HP/GoCode/web/gin/model/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Success bool
	Message string
	Code    int
	Data    interface{}
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Result{
		Success: true,
		Message: constant.SuccessMessage,
		Code:    constant.SuccessCode,
		Data:    data,
	})
}

func Error(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, Result{
		Success: false,
		Message: message,
		Code:    code,
		Data:    "",
	})
	ctx.Abort()
}
