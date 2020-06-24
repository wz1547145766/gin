package user

/*
设置注册相关函数
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB  *gorm.DB
	err error
)

type PostUser struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type User struct {
	Id       int    `gorm:"id" json:"id"`
	Username string `gorm:"username" json:"username" form:"username"`
	Password string `gorm:"password" json:"password" form:"password"`
}

//表单接收创建用户并登陆
func LoginPost(context *gin.Context) {

	var postuser PostUser
	var user User

	//连接数据库
	dbAddr := "root:@/gintest?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dbAddr)
	if err != nil {
		panic("数据库连接失败")
	}
	defer DB.Close()

	//开启自动迁移
	DB.AutoMigrate(&User{})

	//绑定post表单到user
	err := context.ShouldBind(&postuser)

	//字段验证，如果用户名在里面，就返回错误,不在的话就创建新记录
	err = DB.Where("username=?", postuser.Username).First(&user).Error
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"error": "用户名已经存在"})
		return
	} else {
		DB.Create(&User{Username: postuser.Username, Password: postuser.Password})
	}

	//新建cookie
	cookie, err := context.Request.Cookie("username")
	//如果cookie是nil，代表cookie已经存在，已经登录
	if err == nil {
		context.String(http.StatusOK, cookie.Value)
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
