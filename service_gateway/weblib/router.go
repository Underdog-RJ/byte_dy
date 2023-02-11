package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("douyin")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		v1.POST("user/register/", handlers.UserRegister)
		v1.POST("user/login/", handlers.UserLogin)
		v1.GET("user/", handlers.UserInfo)
		// 视频流服务
		v1.GET("feed/", handlers.FeedVideo)
		v1.GET("publish/list/", handlers.VideoList)
		// 需要登录保护
		authed := v1.Group("publish")
		authed.Use(middleware.JWT())
		{
			authed.POST("/action/", handlers.UploadVideo)
		}
	}
	return ginRouter
}
