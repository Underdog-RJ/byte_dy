package core

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"service_video/model"
	"service_video/services"
	"service_video/utils"
)

var VIDEO_OUTPUT_PATH = "F:"

func (s *VideoService) UploadVideo(ctx context.Context, request *services.VideoRequest, response *services.VideoResponse) error {

	// 保存到本地
	currentPath := filepath.Join(VIDEO_OUTPUT_PATH, request.Title)
	ioutil.WriteFile(currentPath, request.Data, 777)

	// 文件扩展名
	ext := filepath.Ext(currentPath)[1:]
	// 获取文件md5
	md5 := utils.GetFileMD5(currentPath)
	var currentVideo = &model.Video{}
	currentVideo.Title = request.Title
	currentVideo.UserId = 1
	if ext == "mp4" {
		currentVideo.VideoStatus = 1
	} else {
		currentVideo.VideoStatus = 0
	}
	currentVideo.VideoMd5 = md5
	currentVideo.VideoExt = ext
	// 上传原始文件
	originalFileName := md5 + "/" + request.Title
	originalFilePath := model.UploadFile("video", currentPath, originalFileName, "video/"+ext)
	currentVideo.OriginFilePath = originalFilePath
	currentVideo.InsertVideo(model.DB)

	response.Code = 200
	return nil

}

func receiveMediaProcessTask(id int64) {
	currentVideo := model.Video{}
	model.DB.First(&currentVideo, id)
	// 获取当前视频的md5
	md5 := currentVideo.VideoMd5
	// 判断当前服务器是否存在此视频，不存在则从minio获取
	currentVideoDir := filepath.Join(VIDEO_OUTPUT_PATH, md5)
	fp := filepath.Join(currentVideoDir, currentVideo.Title)
	_, err := os.Stat(fp)
	if err != nil {
		log.Println("当前文件不存在，从minio获取")
		//todo
	}

	// 获取视频封面
	localCoverName, _ := model.GetSnapshot(fp, VIDEO_OUTPUT_PATH, 1)
	// 封面远程路径
	coverFileName := md5 + "/" + localCoverName
	// 上传视频封面
	coverUrl := model.UploadFile("video", fp, coverFileName, "image/png")
	currentVideo.CoverUrl = coverUrl
	mp4Path := fp
	// todo 视频格式转换,如果视频格式不是mp4转换为mp4格式，
	if currentVideo.VideoExt != "mp4" {

	}

	// 把转换后的视频统一分割为m3u8格式
	m3u8Dir := filepath.Join(currentVideoDir, "hls")
	m3u8Name := filepath.Join(currentVideoDir, "hls", md5+".m3u8")

	// 转换mp4转换为m3u8
	utils.GenerateM3u8(mp4Path, m3u8Name)

	// 上传m3u8文件夹到minio服务器
	dir, _ := os.ReadDir(m3u8Dir)
	for _, fi := range dir {
		cntFilePath := filepath.Join(m3u8Dir, fi.Name())
		remoteFilePath := md5 + "/hls/" + fi.Name()
		model.UploadFile("video", cntFilePath, remoteFilePath, "video/vnd.dlna.mpeg-tts")
	}
	m3u8PlayUrl := "http://159.27.184.52:8888/video/" + md5 + "/hls/" + md5 + ".m3u8"
	currentVideo.PlayUrl = m3u8PlayUrl
	// todo 更新数据库状态

}
