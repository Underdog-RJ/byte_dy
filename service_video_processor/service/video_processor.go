package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"service_video_processor/model"
	"service_video_processor/utils"
)

var VIDEO_OUTPUT_PATH = "F:"

// 处理接收到的消息
func Consumer() {
	ch, err := model.MQ.Channel()
	if err != nil {
		panic(err)
	}
	q, _ := ch.QueueDeclare("video_queue", true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 处于一个监听状态，一致监听我们的生产端的生产，所以这里我们要阻塞主进程
	go func() {
		for d := range msgs {
			var t model.Video
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				log.Fatalf("%v", err)
			}
			ReceiveMediaProcessTask(t.ID)
			fmt.Println("Done")
			_ = d.Ack(false)
		}
	}()
}

func ReceiveMediaProcessTask(id int64) {
	currentVideo := model.Video{}
	model.DB.First(&currentVideo, id)
	// 获取当前视频的md5
	md5 := currentVideo.VideoMd5
	// 判断当前服务器是否存在此视频，不存在则从minio获取
	currentVideoDir := filepath.Join(VIDEO_OUTPUT_PATH, md5)
	// 判断当前文件夹是否存在
	hasDir, _ := utils.HasDir(currentVideoDir)
	// 不存在则创建
	if !hasDir {
		utils.CreateDir(currentVideoDir)
	}
	// 当前文件路径
	fp := filepath.Join(currentVideoDir, currentVideo.Title)
	// 判断当前文件是否存在
	exist, _ := utils.HasDir(fp)
	if !exist {
		// 根据 bucket/md5/title从minio获取，并保存到本地
		remotePath := "/" + md5 + "/" + currentVideo.Title
		localPath := fp
		err := model.FGet(remotePath, localPath)
		if err != nil {
			log.Println("download file exist error")
			panic(err)
		}
	}
	// 构造封面路径
	localCoverPath := filepath.Join(currentVideoDir, md5+".png")
	// 获取视频封面
	err := model.GetSnapshot(fp, localCoverPath, 1)
	if err != nil {
		log.Fatalf("create cover exist error")
	}
	// 封面远程路径
	remoteCoverName := md5 + "/" + md5 + ".png"
	// 上传视频封面
	coverUrl := model.UploadFile("video", localCoverPath, remoteCoverName, "image/png")
	currentVideo.CoverUrl = coverUrl
	mp4Path := fp
	// todo 视频格式转换,如果视频格式不是mp4转换为mp4格式，
	if currentVideo.VideoExt != "mp4" {

	}

	// 把转换后的视频统一分割为m3u8格式
	m3u8Dir := filepath.Join(currentVideoDir, "hls")
	// 判断是否存在m3u8文件夹
	existM3u8Dir, err := utils.HasDir(m3u8Dir)
	if !existM3u8Dir {
		utils.CreateDir(m3u8Dir)
	}

	m3u8Name := filepath.Join(currentVideoDir, "hls", md5+".m3u8")

	// 转换mp4转换为m3u8
	err = utils.GenerateM3u8(mp4Path, m3u8Name)
	if err != nil {
		panic(err)
	}

	// 上传m3u8文件夹到minio服务器
	dir, _ := os.ReadDir(m3u8Dir)
	for _, fi := range dir {
		cntFilePath := filepath.Join(m3u8Dir, fi.Name())
		remoteFilePath := md5 + "/hls/" + fi.Name()
		model.UploadFile("video", cntFilePath, remoteFilePath, "video/vnd.dlna.mpeg-tts")
	}
	m3u8PlayUrl := "http://159.27.184.52:8888/video/" + md5 + "/hls/" + md5 + ".m3u8"
	currentVideo.PlayUrl = m3u8PlayUrl
	// 更新数据库状态 PlayUrl 和 CoverUrl
	model.DB.Save(currentVideo)

}
