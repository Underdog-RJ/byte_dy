package middleware

import (
	"api-gateway/weblib/handlers"
	"errors"
	"service_common/pkg/utils"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		token = c.Query("token")
		if token == "" {
			token = c.PostForm("token")
			if token == "" {
				handlers.PanicIfUserError(errors.New("请先登录"))
			}
		}
		parseToken, err := utils.ParseToken(token)
		if err != nil {
			handlers.PanicIfUserError(errors.New("鉴权失败"))
		}
		c.Set("parseToken", parseToken)
		c.Set("parseUserId", parseToken.Id)
		c.Next()
	}
}
