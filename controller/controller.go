package controller

import (
	"TodoList/dao"
	"TodoList/models"
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)

	c.Redirect(301, "/") //重定向不能用200
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Printf("username: %v\n", username)

	user := dao.Mgr.Login(username)

	if user.Username == "" {
		c.HTML(200, "login.html", "用户名不存在")
		fmt.Printf("用户名不存在")
	} else {
		if user.Password != password {
			c.HTML(200, "login.html", "密码错误")
			fmt.Printf("密码错误")
		} else {
			// c.Redirect(301, "/index.html")
			c.HTML(200, "index.html", nil)
			fmt.Printf("登录成功")
		}

	}
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "user.html", nil)
}

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPosts()
	c.HTML(200, "postIndex.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	tag := c.PostForm("tag")

	post := models.Post{
		Title:   title,
		Content: content,
		Tag:     tag,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(301, "/post_index")
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))

	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
