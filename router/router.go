package router

import (
	"TodoList/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")
	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)
	// e.GET("/index", controller.ListUser)
	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)
	e.GET("/", controller.Index)

	// 博客操作
	e.GET("/post_index", controller.GetPostIndex)
	e.POST("/post", controller.AddPost)
	e.GET("/post", controller.GoAddPost)
	e.GET("/post_detail", controller.PostDetail)

	e.Run()
}
