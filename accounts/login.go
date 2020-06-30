package accounts

/*
设置登录相关函数
*/

import (
	"gin/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err interface{}

//登陆首页
func Login(context *gin.Context) {

	var postUser PostUser
	var user sql.User
	msg := ""

	//连接数据库
	sql.LinkSql()
	defer sql.DB.Close()

	//绑定post表单到user
	err = context.ShouldBind(&postUser)

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
			msg = "登录成功"
		} else {
			err = "账号或密码错误"
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"err": err,
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
