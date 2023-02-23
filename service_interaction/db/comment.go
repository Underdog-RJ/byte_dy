package db

import (
	"errors"
	"interaction/pkg/util"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Id          int64 `gorm:"primaryKey;autoIncrement"`
	UserId      int64
	VideoId     int64
	CommentText string
	CreateTime  time.Time `gorm:"autoCreateTime"`
	IsDel       int8
}

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
	Follow_count   int `gorm:"default:0;"`
	Follower_count int `gorm:"default:0;"`
}

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
	OriginalName   string
}

// InsertComment 插入一条评论
func (c *Comment) InsertComment() (*Comment, error) {
	//数据库中插入一条评论信息
	err := Db.Model(Comment{}).Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

// DeleteComment 删除评论
func (c *Comment) DeleteComment() error {
	//先查询是否有此评论
	result := Db.Model(Comment{}).Where(map[string]interface{}{"id": c.Id, "is_del": util.CommentIsNotDel}).First(&c)
	if result.RowsAffected == 0 { //查询到此评论数量为0则返回无此评论
		return errors.New("delete comment not exist")
	}
	//数据库中删除评论-更新评论状态为-1
	c.IsDel = util.CommentIsDel
	err := Db.Save(&c).Error
	if err != nil {
		return errors.New("delete comment failed")
	}
	return nil
}

func FindUser(userId int64) error {
	user := User{}
	if err := Db.Model(&User{}).First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}
	return nil
}

func FindVideo(videoId int64) error {
	video := Video{}
	if err := Db.Model(&Video{}).First(&video, videoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("视频不存在")
		}
		return err
	}
	return nil
}

func UserInfo(userId int64) *User {
	user := User{}
	Db.Model(&User{}).First(&user, userId)
	return &user
}
