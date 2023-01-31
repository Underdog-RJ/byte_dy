package db

import (
	"errors"
	"gorm.io/gorm"
	"interaction/pkg/util"

	"time"
)

type TbLike struct {
	ID         int64 `gorm:"primaryKey"`
	UserId     int64
	VideoId    int64
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	IsDel      int8
}

func (t TbLike) TableName() string {
	return "likes"
}

// InsertLike 插入点赞数据
func (l *TbLike) InsertLike(db *gorm.DB) error {
	//创建点赞数据，默认为点赞，cancel为0，返回错误结果
	err := db.Model(TbLike{}).Create(l).Error

	//如果有错误结果，返回插入失败
	if err != nil {
		return errors.New("insert data fail")
	}
	return nil
}

// UpdateLike 更新点赞数据 取消点赞或重新点赞
func (t *TbLike) UpdateLike(db *gorm.DB) error {
	err := db.Model(TbLike{}).Where(map[string]interface{}{"user_id": t.UserId, "video_id": t.VideoId}).
		Update("cancel", t.IsDel).Error
	if err != nil {
		return errors.New("update data failed ")
	}
	return nil
}

// GetLikeInfo 获取具体点赞信息
func (t *TbLike) GetLikeInfo() (TbLike, error) {
	//创建一条空like结构体，用来存储查询到的信息
	var likeInfo TbLike
	//根据userid,videoId查询是否有该条信息，如果有，存储在likeInfo,返回查询结果
	err := Db.Model(TbLike{}).Where(map[string]interface{}{"user_id": t.UserId, "video_id": t.VideoId}).
		First(&likeInfo).Error
	if err != nil {
		//查询数据为0，打印"can't find data"，返回空结构体，这时候就应该要考虑是否插入这条数据了
		if "record not found" == err.Error() {
			return TbLike{}, nil
		} else {
			//如果查询数据库失败，返回获取likeInfo信息失败
			return likeInfo, errors.New("get likeInfo failed")
		}
	}
	return likeInfo, nil
}

// GetLikeVideoIdList 根据用户ID查找喜欢列表
func (l *TbLike) GetLikeVideoIdList() ([]int64, error) {
	var likeVideoIdList []int64
	err := Db.Model(TbLike{}).Where(map[string]interface{}{"user_id": l.UserId, "is_del": util.ISLIKE}).
		Pluck("video_id", &likeVideoIdList).Error
	if err != nil {
		//查询数据为0，返回空likeVideoIdList切片，以及返回无错误
		if "record not found" == err.Error() {
			return likeVideoIdList, nil
		} else {
			//如果查询数据库失败，返回获取likeVideoIdList失败
			return likeVideoIdList, errors.New("get likeVideoIdList failed")
		}
	}
	return likeVideoIdList, nil
}
