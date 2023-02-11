package handlers

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"service_common/pkg/utils"
	"service_common/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	ginCtx.JSON(http.StatusOK, gin.H{"status_code": videoResp.GetStatusCode(), "status_msg": videoResp.GetStatusMsg()})

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

// 视频列表
type videolistresponse struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Videolist  []videoinfo `json:"video_list"`
}

// 视频信息
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

// 作者信息
type authorinfo struct {
	Id              int    `json:"id"`
	User_name       string `gorm:"unique"`
	Password_Digest string
	Follow_count    int
	Follower_count  int
	Is_follow       bool `json:"is_follow"`
}
type likes struct {
	User_id  int
	Vider_id int
}

func VideoList(ginCtx *gin.Context) {
	user_idstring := ginCtx.Query("user_id")
	user_id := stringtonum(user_idstring)
	token := ginCtx.Query("token")
	my_userid, _ := utils.ParseToken(token)
	conn, err := gorm.Open("mysql", "root:Zhangzhengxu123.@tcp(159.27.184.52:6033)/ByteQingXun")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	var author authorinfo
	var ress bool
	var foller Follower
	conn.Raw("select id,user_name,follow_count,follower_count from user where id=?", user_id).Scan(&author)
	if err2 := conn.Raw("select follower_id,followee_id from follower where follower_id=?&&followee_id=?", user_id, my_userid.Id).First(&foller).Error; err2 != nil {
		err2 = errors.New("没关注")
		ress = false
	} else {
		ress = true
	}
	author.Is_follow = ress
	log.Println("@@ ", author, " @@")

	var videolist []videoinfo
	conn.Raw("select id,play_url,cover_url,favorite_count,comment_count,title from video where user_id=?", user_id).Scan(&videolist)
	for i := 0; i < len(videolist); i++ {
		videolist[i].Authorinfo = author
		var islike bool
		var like likes
		if err2 := conn.Raw("select user_id,video_id from likes where user_id=?&&video_id=?", user_id, videolist[i].ID).First(&like).Error; err2 != nil {
			err2 = errors.New("不喜欢")
			islike = false
		} else {
			islike = true
		}
		videolist[i].IsFavorite = islike
		//log.Println(videolist[i].IsFavorite, videolist[i].ID, videolist[i].Authorinfo.Id)
	}
	res := videolistresponse{
		StatusCode: 0,
		StatusMsg:  "发布列表",
		Videolist:  videolist,
	}
	ginCtx.JSON(200, res)
	defer conn.Close()
}
