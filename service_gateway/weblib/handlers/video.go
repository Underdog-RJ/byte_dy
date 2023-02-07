package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"service_common/pkg/utils"
	"service_common/services"
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

}
