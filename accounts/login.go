package accounts

/*
设置登录相关函数
*/

import (
	"fmt"
	"gin/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

//登陆首页
func Login(context *gin.Context) {

	var postUser PostUser
	var user sql.User

	var msg string

	//连接数据库
	sql.LinkSql()
	defer sql.DB.Close()

	//绑定post表单到user
	err := context.ShouldBind(&postUser)
	message := context.PostForm("username")
	fmt.Println(postUser)
	fmt.Print(message)

	//查询字段
	err = sql.DB.Where("username = ?", postUser.Username).First(&user).Error
	if err == nil {
		if user.Password == postUser.Password {
			cookie, _ := context.Request.Cookie("username")
			cookie = &http.Cookie{
				Name:  "username",
				Value: postUser.Username,
				Path:  "/",
			}
			http.SetCookie(context.Writer, cookie)
			msg = "success"
		} else {
			context.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"err": err.Error(),
	})
}

//设置清除cookie，退出登录
func Logout(context *gin.Context) {
	cookie, err := context.Request.Cookie("username")
	if err == nil {
		cookie.Value = ""
		cookie.MaxAge = -1
		http.SetCookie(context.Writer, cookie)
	}
}
