package main

import (
	"gin/middleware"
	"gin/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.Use(middleware.Cors())
	//r.Use(islog(true))
	r.LoadHTMLFiles("./index.html")

	shop := r.Group("/index")
	shop.Use(middleware.Islog(true))
	{
		shop.GET("/", func(context *gin.Context) { context.JSON(http.StatusOK, gin.H{"message": "ok"}) })
	}

	r.GET("/", user.Login)
	r.POST("/login", user.LoginPost)
	r.GET("/logout", user.Logout)
	r.Run(":8000")
}
