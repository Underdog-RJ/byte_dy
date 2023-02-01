package handlers

import (
	"api-gateway/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadVideo(ginCtx *gin.Context) {
	file, err := ginCtx.FormFile("data")
	if err != nil {
		fmt.Println("file is null")
		return
	}

	fmt.Println(file.Filename)
	fmt.Println(file.Size)
	var videoReq services.VideoRequest
	videoReq.Title = file.Filename
	PanicIfVideoError(ginCtx.Bind(&videoReq))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.UploadVideo(context.Background(), &videoReq)
	PanicIfVideoError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": videoResp})
}
