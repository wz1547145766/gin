package user

/*
设置登录相关函数
*/

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

//登陆首页
func Login(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"mas": "asd",
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
