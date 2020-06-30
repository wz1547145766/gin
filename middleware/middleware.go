package middleware

import (
	"net/http"
	"regexp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	if gin.Mode() == gin.ReleaseMode {
		// 生产环境需要配置跨域域名，否则403
		config.AllowOrigins = []string{"http://www.example.com"}
	} else {
		// 测试环境下模糊匹配本地开头的请求
		config.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	config.AllowCredentials = true
	return cors.New(config)
}

//检测登录状态中间件
func Islog(check bool) gin.HandlerFunc {

	return func(context *gin.Context) {
		cookie, err := context.Request.Cookie("username")
		if err == nil {
			context.JSON(http.StatusOK, gin.H{
				"cookie": cookie.Value,
			})
			context.Next()
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "没有登录,此操作需要登录才能执行",
				"err": err,
			})
			context.Abort()
		}
	}
}
