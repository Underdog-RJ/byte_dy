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
		// 视频流服务
		v1.GET("feed/", handlers.FeedVideo)

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 获取用户信息
			authed.GET("user/", handlers.UserInfo)
			// 用户投稿
			authed.POST("publish/action/", handlers.UploadVideo)
			// 用户列表信息
			authed.GET("publish/list/", handlers.VideoList)
			authed.POST("/relation/action", handlers.Relation.RelationAction)
			authed.GET("/relation/follow/list", handlers.Relation.RelationFollowList)
			authed.GET("/relation/follower/list", handlers.Relation.RelationFollowerList)
			authed.GET("/relation/friend/list", handlers.Relation.RelationFriendList)
			authed.POST("/message/action", handlers.Relation.MessageAction)
			authed.GET("/message/chat", handlers.Relation.MessageChat)
		}
	}
	return ginRouter
}
