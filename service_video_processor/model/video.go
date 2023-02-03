package model

import (
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
