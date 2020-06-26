package sever

import (
	"gin/accounts"
	"gin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

//API接口
func Routers() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.Cors())

	//主页
	r.GET("/")

	//待定,暂时为检验是否登录
	shop := r.Group("/index")
	shop.Use(middleware.Islog(true))
	{
		shop.GET("/", func(context *gin.Context) { context.JSON(http.StatusOK, gin.H{"message": "ok"}) })
	}

	//账号相关
	account := r.Group("accounts")
	{
		account.POST("/login", accounts.Login)
		account.POST("/reg", accounts.LoginPost)
		account.GET("/logout", accounts.Logout)
	}

	return r
}
