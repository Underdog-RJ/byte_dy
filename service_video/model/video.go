package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Video struct {
	ID             int64 `gorm:"primaryKey"`
	UserId         int64
	PlayUrl        string
	CoverUrl       string
	FavoriteCount  int64
	CommentCount   int64
	PublishTime    time.Time
	Title          string
	VideoStatus    int64
	VideoSize      int64
	VideoExt       string
	VideoMd5       string
	OriginFilePath string
}

func (l *Video) InsertVideo(db *gorm.DB) error {
	//创建点赞数据，默认为点赞，cancel为0，返回错误结果
	err := db.Model(Video{}).Create(l).Error

	//如果有错误结果，返回插入失败
	if err != nil {
		return errors.New("insert data fail")
	}
	return nil
}
