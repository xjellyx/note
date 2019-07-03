package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s", name)
	})
	router.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		msg := name + "is" + action
		ctx.JSON(http.StatusOK, gin.H{
			"data": msg,
		})
	})
	router.POST("/formPost", func(ctx *gin.Context) {
		msg := ctx.PostForm("message")
		test := ctx.PostForm("test")
		data := make(map[string]interface{})
		data["test"] = test
		nick := ctx.DefaultPostForm("nick", "anonymous")
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": msg,
			"nick":    nick,
			"data":    data,
		})
	})
	//router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//router.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Main website",
	//	})
	//}

	router.Run(":8080")
	router.Run()
}
