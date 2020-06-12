package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string  `json:"username" form:"username"`
	Password string  `json:"password" form:"password"`
}


//登陆首页
func login (context *gin.Context) {

	context.HTML(http.StatusOK, "index.html", gin.H{
		"message":"ok",
	})
}

//表单接收创建
func loginPost (context *gin.Context) {

	var user User

	err := context.ShouldBind(&user)  //绑定post表单到user

	cookie, err := context.Request.Cookie("username")
	if err == nil {
		context.String(http.StatusOK, cookie.Value)
	} else {
		cookie = &http.Cookie{
			Name:  "username",
			Value: user.Username,
		}
		http.SetCookie(context.Writer, cookie)

		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": user.Username,
			"cookie":   cookie.Value,
		})
	}
}


//设置清除cookie，退出登录
func logout(context *gin.Context){
	cookie, err := context.Request.Cookie("username")
	if err == nil{
		cookie.Value = ""
		cookie.MaxAge = -1
		http.SetCookie(context.Writer, cookie)
	}
}


//检测登录状态中间件
func islog(check bool) gin.HandlerFunc {

	return func(context * gin.Context){
		cookie, err := context.Request.Cookie("username")
		if err == nil {
			context.JSON(http.StatusOK, gin.H{
				"cookie": cookie.Value,
			})
			context.Next()
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "没有登录",
			})
			context.Abort()
		}
	}
}



func main() {

	r := gin.Default()
	//r.Use(islog(true))
	r.LoadHTMLFiles("./index.html")

	shop := r.Group("/index")
	shop.Use(islog(true))
	{
		shop.GET("/", func(context *gin.Context) {context.JSON(http.StatusOK, gin.H{"message":"ok"})})
	}


	r.GET("/", login)
	r.POST("/login", loginPost)

	r.GET("/logout", logout)

	r.Run()
}
