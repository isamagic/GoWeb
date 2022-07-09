package main

import (
	"learning/GoWeb/engine"
	"net/http"
)

func main() {
	r := engine.New()

	r.GET("/", func(ctx *engine.Context) {
		ctx.HTML(http.StatusOK, "<h1>hello world<h1>")
	})

	// 格式：/hello?name=xxx
	r.GET("/hello", func(ctx *engine.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	// 格式：/hello/geek
	r.GET("/hello/:name", func(c *engine.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	// 格式：/assets/css/geek.css
	r.GET("/assets/*filepath", func(c *engine.Context) {
		c.JSON(http.StatusOK, engine.H{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(ctx *engine.Context) {
		ctx.JSON(http.StatusOK, engine.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	r.Run(":9999")
}
