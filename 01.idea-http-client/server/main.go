package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/guuid"
)

var token string

func main() {
	s := g.Server()
	group := s.Group("/api")
	// 默认路径
	// GET带参数
	group.GET("/get", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
		r.Response.Writeln("name:", r.GetString("name"))
	})
	// POST KV
	group.POST("/post/kv", func(r *ghttp.Request) {
		r.Response.Writeln("func:test")
		r.Response.Writeln("name:", r.GetString("name"))
		r.Response.Writeln("age:", r.GetInt("age"))
	})
	// POST JSON
	group.POST("/post/json", func(r *ghttp.Request) {
		r.Response.Writeln("func:test2")
		r.Response.Writeln("name:", r.GetString("name"))
		r.Response.Writeln("age:", r.GetString("age"))

		h := r.Header
		r.Response.Writeln("referer:", h.Get("referer"))
		r.Response.Writeln("cookie:", h.Get("cookie"))
		r.Response.Writeln(r.Cookie.Map())
	})

	// 模拟登陆
	system := s.Group("/system")
	// 登陆接口
	system.POST("/login", func(r *ghttp.Request) {
		if "admin" == r.GetString("username") &&
			"123456" == r.GetString("password") {
			token = guuid.New().String()
			r.Response.WriteJson(g.Map{
				"code": 0,
				"data": token,
			})
			r.Exit()
		}
		r.Response.WriteJson(g.Map{
			"code": -1,
			"data": "",
		})
	})
	// 获取用户信息
	system.POST("/user/info", func(r *ghttp.Request) {
		if token != r.Header.Get("token") || token == "" {
			r.Response.WriteJson(g.Map{
				"code": -1,
				"data": "",
			})
			r.Exit()
		}

		// 返回用户信息
		r.Response.WriteJson(g.Map{
			"code": 0,
			"data": "zhangsan",
		})
	})
	// 获取用户年龄
	system.POST("/user/age", func(r *ghttp.Request) {
		if token != r.Header.Get("token") || token == "" {
			r.Response.WriteJson(g.Map{
				"code": -1,
				"data": "",
			})
			r.Exit()
		}

		// 返回用户信息
		r.Response.WriteJson(g.Map{
			"code": 0,
			"data": 11,
		})
	})

	s.SetPort(80)
	s.Run()
}
