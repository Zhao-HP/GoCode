package router

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRoute() {

	fmt.Println("初始化路由")

	s := g.Server()

	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello GoFrame")
	})

}

func Test() {
	fmt.Println("bbbbbb")
}
