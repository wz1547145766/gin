package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	err interface{}
)

//用户表
type User struct {
	Id       int    `gorm:"id" json:"id"`
	Username string `gorm:"username" json:"username" form:"username"`
	Password string `gorm:"password" json:"password" form:"password"`
}

//新闻表
type News struct {
	Id      int    `gorm:"id" json:"id"`
	Title   string `gorm:"title" json:"title"`
	Content string `gorm:"content" json:"content"`
}

//连接数据库
func LinkSql() {
	dbAddr := "root:@/gorm?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dbAddr)
	if err != nil {
		panic("数据库连接失败")
	}
}
