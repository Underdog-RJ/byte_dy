package db

import (
	"errors"
	"gorm.io/gorm"
	"interaction/pkg/util"
	"time"
)

type TbComment struct {
	Id          int64 `gorm:"primaryKey"`
	UserId      int64
	VideoId     int64
	CommentText string
	CreateTime  time.Time `gorm:"autoCreateTime"`
	IsDel       int8
}

func (c TbComment) TableName() string {
	return "comment"
}

// InsertComment 插入一条评论
func (c *TbComment) InsertComment(db *gorm.DB) (*TbComment, error) {
	//数据库中插入一条评论信息
	err := Db.Model(TbComment{}).Create(&c).Error
	if err != nil {
		return &TbComment{}, errors.New("create comment failed")
	}
	return c, nil
}

// DeleteComment 删除评论
func (c *TbComment) DeleteComment(db *gorm.DB) error {
	//先查询是否有此评论
	result := Db.Model(TbComment{}).Where(map[string]interface{}{"id": c.Id, "cancel": util.CommentIsNotDel}).First(&c)
	if result.RowsAffected == 0 { //查询到此评论数量为0则返回无此评论
		return errors.New("del comment is not exist")
	}
	//数据库中删除评论-更新评论状态为-1
	c.IsDel = util.CommentIsDel
	err := db.Save(&c).Error
	if err != nil {
		return errors.New("del comment failed")
	}
	return nil
}
