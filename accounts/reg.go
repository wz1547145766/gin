package accounts

/*
设置注册相关函数
*/

import (
	"gin/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PostUser struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

//表单接收创建用户并登陆
func LoginPost(context *gin.Context) {

	var postuser PostUser
	var user sql.User

	//连接数据库
	sql.LinkSql()
	defer sql.DB.Close()

	//开启自动迁移
	sql.DB.AutoMigrate(&sql.User{})

	//绑定post表单到user
	err := context.ShouldBind(&postuser)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"err": err})
	}

	//字段验证，如果用户名存在，就返回错误,不在的话就创建新记录
	err = sql.DB.Where("username=?", postuser.Username).First(&user).Error
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"error": "用户名已经存在"})
		return
	} else {
		sql.DB.Create(&sql.User{Username: postuser.Username, Password: postuser.Password})
	}

	//新建cookie
	cookie, err := context.Request.Cookie("username")
	//如果cookie是nil，代表cookie已经存在，已经登录
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"error": "您已经登录"})
		return
	} else {
		cookie = &http.Cookie{
			Name:  "username",
			Value: postuser.Username,
		}
		http.SetCookie(context.Writer, cookie)

		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": postuser.Username,
			"cookie":   cookie.Value,
		})
	}
}
