package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UploadVideo(ginCtx *gin.Context) {
	formFile, err := ginCtx.FormFile("data")

	if err != nil {
		fmt.Println("file is null")
		return
	}
	token := ginCtx.PostForm("token")
	if token == "" {
		fmt.Println("权限不足")
		return
	}

	// 获取用户id
	parseToken, _ := utils.ParseToken(token)
	println(int(parseToken.Id))
	userId := strconv.FormatInt(int64(int(parseToken.Id)), 10)

	file, err := formFile.Open()

	bytes, err := ioutil.ReadAll(file)

	// 关闭文件流
	defer file.Close()

	if err != nil {
		fmt.Println("打开文件出错")
	}

	var videoReq services.VideoRequest
	videoReq.Title = formFile.Filename
	videoReq.Token = userId
	videoReq.Data = bytes

	//PanicIfVideoError(ginCtx.Bind(&videoReq))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.UploadVideo(context.Background(), &videoReq)
	PanicIfVideoError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": videoResp})
}
