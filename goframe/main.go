package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-code/config"
	"goframe-code/router"
)

func main() {
	// 初始化路由
	router.InitRoute()

	server := g.Server()
	server.SetPort(config.GLOBAL_CONF.Server.HttpPort)

	server.Run()

}
