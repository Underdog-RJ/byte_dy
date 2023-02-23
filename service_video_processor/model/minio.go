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
	"os"
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

func GetSnapshot(videoPath, filename string, frameNum int) (err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).WithOutput(buf, os.Stdout).Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}
	err = imaging.Save(img, filename)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}
	return nil

}

func UploadFile(bucketName, filepath, filename, contentType string) string {
	ctx := context.Background()

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("current file path is null")
	}

	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalf("current file content is null")
	}

	// 判断bucket是否存在
	exists, _ := minioClient.BucketExists(ctx, bucketName)

	if !exists {
		minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}

	n, err := minioClient.PutObject(ctx, bucketName, filename, file, fileStat.Size(), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalf("current bucket exist error")
	}

	log.Println("Successfully uploaded bytes: ", n)
	fileUrl := "http://159.27.184.52:8888/video/" + filename
	log.Println("generate fileUrl success:" + fileUrl)
	return fileUrl
}

func FGet(remotePath, localPath string) error {
	ctx := context.Background()
	err := minioClient.FGetObject(ctx, "video", remotePath, localPath, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	return err
}
