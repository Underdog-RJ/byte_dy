package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"service_common/pkg/utils"
	"service_common/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadVideo(ginCtx *gin.Context) {
	formFile, err := ginCtx.FormFile("data")
	if err != nil {
		log.Println("file is null")
		return
	}
	token := ginCtx.PostForm("token")
	if token == "" {
		log.Println("权限不足")
		return
	}
	title := ginCtx.PostForm("title")

	// 获取用户id
	parseToken, _ := utils.ParseToken(token)

	userId := parseToken.Id

	file, err := formFile.Open()
	bytes, err := ioutil.ReadAll(file)

	// 关闭文件流
	defer file.Close()

	if err != nil {
		fmt.Println("打开文件出错")
	}

	var videoReq services.VideoRequest
	videoReq.OriginalName = formFile.Filename
	videoReq.UserId = int64(userId)
	videoReq.Data = bytes
	videoReq.Title = title

	//PanicIfVideoError(ginCtx.Bind(&videoReq))

	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.UploadVideo(context.Background(), &videoReq)
	PanicIfVideoError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": videoResp})

}

func FeedVideo(ginCtx *gin.Context) {
	token := ginCtx.PostForm("token")
	if token == "" {
		log.Println("未登录")
	}
	latest_time := ginCtx.DefaultQuery("latest_time", "0")
	videoService := ginCtx.Keys["videoService"].(services.VideoService)
	var feedReq services.DouyinFeedRequest
	feedReq.Token = token
	i, _ := strconv.ParseInt(latest_time, 10, 64)
	feedReq.LatestTime = i
	videoResp, err := videoService.FeedVideo(context.Background(), &feedReq)

	PanicIfVideoError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"status_code": videoResp.GetStatusCode(), "status_msg": videoResp.GetStatusMsg(), "video_list": videoResp.GetVideoList(), "next_time": videoResp.NextTime})

}

type videolistresponse struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Videolist  []videoinfo `json:"video_list"`
}
type videoinfo struct {
	ID            int        `json:"id"`
	Authorinfo    authorinfo `json:"author"`
	PlayURL       string     `json:"play_url"`
	CoverURL      string     `json:"cover_url"`
	FavoriteCount int        `json:"favorite_count"`
	CommentCount  int        `json:"comment_count"`
	IsFavorite    bool       `json:"is_favorite"`
	Title         string     `json:"title"`
}
type authorinfo struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func VideoList(ginCtx *gin.Context) {
	user_id := ginCtx.Query("user_id")
	log.Panicln(user_id)
	token := ginCtx.Query("token")
	my_userid, _ := utils.ParseToken(token)
	userinfo := authorinfo{int(my_userid.Id), "user", 1, 1, true}
	var videolist []videoinfo
	video := videoinfo{
		ID:            0,
		Authorinfo:    userinfo,
		PlayURL:       "",
		CoverURL:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "",
	}
	videolist = append(videolist, video)
	res := videolistresponse{
		StatusCode: 0,
		StatusMsg:  "发布列表",
		Videolist:  videolist,
	}
	ginCtx.JSON(200, res)
}
