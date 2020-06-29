package news

import (
	"fmt"
	"gin/sql"

	"github.com/gin-gonic/gin"
)

var err interface{}

//新闻表单提交
type News struct {
	Title   string `gorm:"title" json:"title"`
	Content string `gorm:"content" json:"content"`
}

func Shownews(context *gin.Context) {

	//开启自动迁移
	sql.DB.AutoMigrate(&sql.User{})

	sql.LinkSql()
	err = sql.DB.Find(&sql.News{}).Error
	if err == nil {
		fmt.Println(&sql.News{})
	}

}
