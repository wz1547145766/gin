package sever

import (
	"gin/accounts"
	"gin/middleware"
	"gin/news"
	"net/http"

	"github.com/gin-gonic/gin"
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

	//账号
	account := r.Group("accounts")
	{
		account.POST("/login", accounts.Login)
		account.POST("/reg", accounts.LoginPost)
		account.GET("/logout", accounts.Logout)
	}

	//新闻
	new := r.Group("news")
	new.Use(middleware.Islog(true))
	{
		//查看所有新闻
		new.GET("/show", news.Shownews)
		//查看当前新闻
		new.GET("/current/:id", news.Currentnews)
		//删除新闻
		new.GET("/delnews/:id", news.Delnews)
		//更新新闻
		new.POST("/update", news.Updatanews)
		//增加新闻
		new.POST("/add", news.Addnews)
	}

	return r
}
