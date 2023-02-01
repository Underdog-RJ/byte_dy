package model

import (
	"bytes"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

var minioClient *minio.Client

func Minio(minioMap map[string]string) {
	//根据配置文件连接minio server
	client, err := minio.New(minioMap["MinioUrl"]+":"+minioMap["MinioPort"], &minio.Options{
		Creds:  credentials.NewStaticV4(minioMap["MinioAccessKey"], minioMap["MinioSecretKey"], ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("conetct minio server fail %s url %s ", err.Error(), minioMap["minioUrl"]+":"+minioMap["minioPort"])
	}
	minioClient = client
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).WithOutput(buf, os.Stdout).Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}

func UploadFile(bucketName, filepath, filename, fileExt string, isVideo bool) string {
	ctx := context.Background()
	name, err2 := GetSnapshot(filename, "test", 1)
	if err2 != nil {
		fmt.Println(err2)
	}
	println(name)
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("current file path is null")
	}

	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalf("current file content is null")
	}
	// 判断ContentType类型
	var contentType string
	if isVideo {
		contentType = "video/" + fileExt
	} else {
		contentType = "image/" + fileExt
	}
	// 判断bucket是否存在
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("current bucket exist error")

	}
	if !exists {
		minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}

	n, err := minioClient.PutObject(ctx, bucketName, filename, file, fileStat.Size(), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalf("current bucket exist error")
	}
	log.Println("Successfully uploaded bytes: ", n)
	fileUrl := GetFileUrl(bucketName, filename, fileExt, isVideo)
	log.Println("generate fileUrl success:" + fileUrl)
	return fileUrl
}

func GetFileUrl(bucketName, filename, fileExt string, isVideo bool) string {
	ctx := context.Background()
	reqParams := make(url.Values)
	//reqParams	url.Values	额外的响应头，支持
	//response-expires，  到期时间
	//response-content-type，  响应内容的类型
	//response-cache-control，   缓存控制
	// response-content-disposition。  内容处理 例如，附件形式强制下载"attachment; filename=\"chat.mp4\""
	// reqParams.Set("response-content-type", "attachment; filename=\"chat.mp4\"")
	// 判断ContentType类型
	var contentType string
	if isVideo {
		contentType = "video/" + fileExt
	} else {
		contentType = "image/" + fileExt
	}
	reqParams.Set("response-content-disposition", contentType)

	presignedURL, err := minioClient.PresignedGetObject(ctx, bucketName, filename, time.Hour*24, reqParams)
	if err != nil {
		log.Fatal("gengerate url failed...")
	}

	return presignedURL.String()

}
