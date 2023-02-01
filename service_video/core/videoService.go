package core

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"service_video/model"
	"service_video/services"
)

func (s *VideoService) UploadVideo(ctx context.Context, request *services.VideoRequest, response *services.VideoResponse) error {
	response.Code = 200
	dir, err := os.Getwd()
	if err != nil {
		log.Println("file exist error")
	}
	fp := dir + "\\" + request.Title
	ext := filepath.Ext(fp)[1:]
	fileUrl := model.UploadFile("video", fp, request.Title, ext, true)
	fmt.Println(fileUrl)
	return nil

}
