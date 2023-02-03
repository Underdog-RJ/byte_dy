package db

import "time"

type TbComment struct {
	Id          int64 `gorm:"primaryKey"`
	UserId      int64
	VideoId     int64
	CommentText string
	CreateTime  time.Time `gorm:"autoCreateTime"`
	IsDel       int8
}
