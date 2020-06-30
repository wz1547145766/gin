package news

import (
	"fmt"
	"gin/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var err interface{}

//新闻表单提交
type PostNews struct {
	Id      int    `gorm:"id" json:"id" form:"id"`
	Title   string `gorm:"title" json:"title" form:"title"`
	Content string `gorm:"content" json:"content" form:"content"`
}

//新添加新闻
func Addnews(context *gin.Context) {

	var postnews PostNews

	//连接数据库
	sql.LinkSql()
	defer sql.DB.Close()

	//开启自动迁移
	sql.DB.AutoMigrate(&sql.News{})

	//绑定用户提交表单
	err = context.ShouldBind(&postnews)
	if err == nil {
		//新建记录
		sql.DB.Create(&sql.News{Title: postnews.Title, Content: postnews.Content})
	} else {
		context.JSON(http.StatusOK, gin.H{"err": err})
	}

	context.JSON(http.StatusOK, gin.H{
		"msg": "新闻添加成功",
	})
}

//删除当期新闻
func Delnews(context *gin.Context) {

	var new sql.News

	sql.LinkSql()
	defer sql.DB.Close()

	newsId := context.Param("id")
	err = sql.DB.Where("id=?", newsId).First(&new).Error
	if err == nil {
		sql.DB.Delete(&new)
	} else {
		context.JSON(http.StatusOK, gin.H{"err": err})
	}
}

//修改当前新闻
func Updatanews(context *gin.Context) {

	var new sql.News
	var postnews PostNews

	sql.LinkSql()
	defer sql.DB.Close()

	//开启自动迁移
	sql.DB.AutoMigrate(&sql.News{})

	err = context.ShouldBind(&postnews)
	if err == nil {
		fmt.Println("1")
		err = sql.DB.Where("id=?", postnews.Id).First(&new).Error
		if err == nil {
			fmt.Println("2")
			new.Title = postnews.Title
			new.Content = postnews.Content
			sql.DB.Save(&new)
		} else {
			fmt.Println("3")
			context.JSON(http.StatusOK, gin.H{"err": err})
		}
	} else {
		fmt.Println("4")
		context.JSON(http.StatusOK, gin.H{"err": err})
	}
}

//查看所有新闻
func Shownews(context *gin.Context) {

	var news []sql.News

	sql.LinkSql()
	defer sql.DB.Close()

	//查询所有新闻，没有新闻返回错误
	err = sql.DB.Find(&news).Error
	if err == nil {
		fmt.Println(&sql.News{})
	} else {
		context.JSON(http.StatusOK, gin.H{"err": err})
	}

	context.JSON(http.StatusOK, gin.H{
		"news": news,
	})

}

//查看当前新闻
func Currentnews(context *gin.Context) {

	var new sql.News

	sql.LinkSql()
	defer sql.DB.Close()

	newsid := context.Param("id")

	err = sql.DB.Where("id = ?", newsid).First(&new).Error
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"msg": new})
	}

}
