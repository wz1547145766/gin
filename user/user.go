package user

/*
设置注册相关函数
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

//表单接收创建
func LoginPost(context *gin.Context) {

	var user User

	err := context.ShouldBind(&user) //绑定post表单到user

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
