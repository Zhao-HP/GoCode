package main

import (
	"github.com/Zhao-HP/GoCode/web/gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router.Init(app)
	_ = app.Run(":9006")
}
