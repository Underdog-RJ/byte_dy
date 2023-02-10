package handlers

import (
	"context"
	"log"
	"net/http"
	"service_common/pkg/utils"
	"service_common/services"

	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
)

// 用户注册
func UserRegister(ginCtx *gin.Context) {
	defer UserPanicHandler(ginCtx)
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.ID))
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.Code,
		"status_msg":  "注册成功",
		"user_id":     userResp.ID,
		"token":       token,
	})
}

// 用户登录
func UserLogin(ginCtx *gin.Context) {
	defer UserPanicHandler(ginCtx)
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.ID))
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": userResp.Code,
		"status_msg":  "登录成功",
		"user_id":     userResp.ID,
		"token":       token,
	})
}

type user struct {
	Id             int    `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	Follow_count   int    `db:"follow_count" json:"follow_count"`
	Follower_count int    `db:"follower_count" json:"follower_count"`
	Is_follow      bool   `db:"is_follow" json:"is_follow"`
	Token          string `db:"token"`
}

// 用户信息
func UserInfo(ginCtx *gin.Context) {
	// conn, err := gorm.Open("mysql", "root:Zhangzhengxu123.@tcp(159.27.184.52:6033)/ByteQingXun")
	// if err != nil {
	// 	fmt.Println("gorm.Open err:", err)
	// 	return
	// }
	// defer conn.Close()
	// var userinfo user
	// conn.First(&userinfo)
	// log.Println(userinfo)
	defer UserPanicHandler(ginCtx)
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserInfo(context.Background(), &userReq)
	PanicIfUserError(err)
	//token, err := utils.GenerateToken(uint(userResp.ID))
	user_id := ginCtx.Query("user_id")
	token := ginCtx.Query("token")
	my_userid, _ := utils.ParseToken(token)
	log.Println(user_id)
	log.Println(token)
	log.Println(my_userid.Id)
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "用户信息",
		"user": map[string]interface{}{
			"id":             userResp.ID,
			"name":           "user",
			"follow_count":   userResp.ID,
			"follower_count": userResp.ID,
			"is_follow":      true,
		},
	})

}
