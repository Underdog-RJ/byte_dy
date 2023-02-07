package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"path/filepath"
	"service_common/services"
	"service_video/model"
	"service_video/utils"
	"time"
)

var VIDEO_OUTPUT_PATH = "F:"

func (s *VideoService) UploadVideo(ctx context.Context, request *services.VideoRequest, response *services.VideoResponse) error {

	// 保存到本地
	currentPath := filepath.Join(VIDEO_OUTPUT_PATH, request.OriginalName)
	ioutil.WriteFile(currentPath, request.Data, 777)

	// 文件扩展名
	ext := filepath.Ext(currentPath)[1:]
	// 获取文件md5
	md5 := utils.GetFileMD5(currentPath)
	var currentVideo = &model.Video{}
	currentVideo.OriginalName = request.OriginalName
	currentVideo.UserId = request.UserId
	if ext == "mp4" {
		currentVideo.VideoStatus = 1
	} else {
		currentVideo.VideoStatus = 0
	}
	currentVideo.VideoMd5 = md5
	currentVideo.VideoExt = ext
	// 上传原始文件
	originalFileName := md5 + "/" + request.OriginalName
	originalFilePath := model.UploadFile("video", currentPath, originalFileName, "video/"+ext)
	currentVideo.OriginFilePath = originalFilePath
	currentVideo.PublishTime = time.Now()
	currentVideo.VideoSize = int64(len(request.Data))
	currentVideo.Title = request.Title
	currentVideo.InsertVideo(model.DB)

	// 发送rabbitMQ消息，处理视频
	err := Producer(currentVideo)
	if err != nil {
		// todo 重试操作
		log.Fatalf("video send rabbitmq error")
	}

	response.StatusCode = 200
	response.StatusMsg = "upload video success"

	return nil

}

// 往rabbitMQ发送消息，处理当前视频
func Producer(videoVo *model.Video) error {
	ch, err := model.MQ.Channel()
	if err != nil {
		err = errors.New("rabbitMQ channel err:" + err.Error())
	}
	q, _ := ch.QueueDeclare("video_queue", true, false, false, false, nil)

	body, _ := json.Marshal(videoVo) // title，content
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		err = errors.New("rabbitMQ publish err:" + err.Error())
	}
	return nil
}
