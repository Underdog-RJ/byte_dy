package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"service_common/pkg/utils"
	"service_common/services"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 用户注册
func UserRegister(ginCtx *gin.Context) {
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

type User struct {
	gorm.Model
	User_name       string `gorm:"unique"`
	Password_Digest string
	Follow_count    int
	Follower_count  int
}
type Follower struct {
	Follower_id int
	Followee_id int
}

func stringtonum(s string) int {
	var a int
	for i := 0; i < len(s); i++ {
		a = a*10 + int(s[i]) - '0'
	}
	return a
}

// 用户信息
func UserInfo(ginCtx *gin.Context) {
	conn, err := gorm.Open("mysql", "root:Zhangzhengxu123.@tcp(159.27.184.52:6033)/ByteQingXun")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}

	var userinfo User
	var foller Follower
	defer UserPanicHandler(ginCtx)
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	PanicIfUserError(err)
	user_idstring := ginCtx.Query("user_id")
	user_id := stringtonum(user_idstring)
	token := ginCtx.Query("token")
	my_userid, _ := utils.ParseToken(token)
	var res bool
	conn.Raw("select id,user_name,follow_count,follower_count from user where id=?", user_id).Scan(&userinfo)
	log.Println("??? ", userinfo.User_name, userinfo.ID, user_id, userinfo.Follow_count, userinfo.Follower_count, my_userid.Id, "???")
	if err2 := conn.Raw("select follower_id,followee_id from follower where follower_id=?&&followee_id=?", user_id, my_userid.Id).First(&foller).Error; err2 != nil {
		err2 = errors.New("没关注")
		res = false
		log.Println("? ", err2, foller, " ? ")
	} else {
		log.Println("? 关注了", foller, " ?")
		res = true
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "用户信息",
		"user": map[string]interface{}{
			"id":             user_id,
			"name":           userinfo.User_name,
			"follow_count":   userinfo.Follow_count,
			"follower_count": userinfo.Follower_count,
			"is_follow":      res,
		},
	})
	defer conn.Close()
}
