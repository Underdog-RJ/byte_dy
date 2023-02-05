package handle

import (
	"context"
	"interaction/db"
	"interaction/middleware/rabbitmq"
	"interaction/middleware/redis"
	"interaction/pkg/util"
	"interaction/service"
	"strconv"
	"strings"
	"time"
)

type LikeService struct{}

func (l *LikeService) IsLike(ctx context.Context, req *service.IsLikeRequest) (*service.IsLikeResponse, error) {
	strUserId := util.LikeUserKey + strconv.FormatInt(req.UserId, 10)
	strVideoId := util.LikeVideoKey + strconv.FormatInt(req.VideoId, 10)
	resp := new(service.IsLikeResponse)
	resp.Code = util.Success

	return resp, nil
}

// 执行点赞或取消点赞操作
func (l *LikeService) LikeAction(ctx context.Context, req *service.LikeActionRequest) (*service.LikeActionResponse, error) {

	strUserId := util.LikeUserKey + strconv.FormatInt(req.UserId, 10)
	strVideoId := util.LikeVideoKey + strconv.FormatInt(req.VideoId, 10)
	resp := new(service.LikeActionResponse)
	resp.Code = util.Success

	var likeDao *db.TbLike

	sb := strings.Builder{}
	sb.WriteString(strconv.FormatInt(req.UserId, 10))
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatInt(req.VideoId, 10))

	// 判断操作类型 是点赞还是不点赞
	if req.ActionType == util.ISLIKE {
		//查询Redis RdbLike(key:strUserId)是否已经加载过此信息
		if n, err := redis.RdbLike.Exists(ctx, strUserId).Result(); n > 0 {
			//如果有问题，说明查询redis失败,返回错误信息
			if err != nil {
				resp.Code = util.Error
				return resp, err
			}
			if _, err1 := redis.RdbLike.SAdd(ctx, strUserId, req.VideoId).Result(); err1 != nil {
				resp.Code = util.Error
				return resp, err1
			}
		} else {
			likeDao = &db.TbLike{UserId: req.UserId, VideoId: req.VideoId, IsDel: int8(req.ActionType)}
			if videoList, err1 := likeDao.GetLikeVideoIdList(); err1 != nil {
				// todo 打印日志
			} else {
				videoList = append(videoList, req.VideoId)
				go addRelationToLike(strUserId, videoList)
			}
		}
		//查询Redis RdbLike(key:strVideoId)是否已经加载过此信息
		if n, err := redis.RdbLike.Exists(ctx, strVideoId).Result(); n > 0 {
			if err != nil {
				resp.Code = util.Error
				return resp, err
			}
			if _, err1 := redis.RdbLike.SAdd(ctx, strVideoId, req.UserId).Result(); err1 != nil {
				resp.Code = util.Error
				return resp, err1
			}
		} else {
			if likeDao == nil {
				likeDao = &db.TbLike{UserId: req.UserId, VideoId: req.VideoId, IsDel: int8(req.ActionType)}
			}
			if userList, err1 := likeDao.GetLikeUserIdList(); err1 != nil {
				// todo 打印日志
			} else {
				userList = append(userList, req.UserId)
				go addRelationToLike(strVideoId, userList)
			}
		}
		// 向消息队列发送消息
		rabbitmq.RmqLike.Publish(sb.String())

	} else {
		//查询Redis RdbLike(key:strUserId)是否已经加载过此信息
		if n, err := redis.RdbLike.Exists(ctx, strUserId).Result(); n > 0 {
			//如果有问题，说明查询redis失败,返回错误信息
			if err != nil {
				resp.Code = util.Error
				return resp, err
			}
			if _, err1 := redis.RdbLike.SRem(ctx, strUserId, req.VideoId).Result(); err1 != nil {
				resp.Code = util.Error
				return resp, err1
			}
		} else {
			likeDao = &db.TbLike{UserId: req.UserId, VideoId: req.VideoId, IsDel: int8(req.ActionType)}
			if videoList, err1 := likeDao.GetLikeVideoIdList(); err1 != nil {
				// todo 打印日志
			} else {
				for i := 0; i < len(videoList); i++ {
					if videoList[i] == req.VideoId {
						videoList = append(videoList[:i], videoList[i+1:]...)
					}
				}
				go addRelationToLike(strUserId, videoList)
			}
		}
		//查询Redis RdbLike(key:strVideoId)是否已经加载过此信息
		if n, err := redis.RdbLike.Exists(ctx, strVideoId).Result(); n > 0 {
			if err != nil {
				resp.Code = util.Error
				return resp, err
			}
			if _, err1 := redis.RdbLike.SRem(ctx, strVideoId, req.UserId).Result(); err1 != nil {
				resp.Code = util.Error
				return resp, err1
			}
		} else {
			if likeDao == nil {
				likeDao = &db.TbLike{UserId: req.UserId, VideoId: req.VideoId, IsDel: int8(req.ActionType)}
			}
			if userList, err1 := likeDao.GetLikeUserIdList(); err1 != nil {
				// todo 打印日志
			} else {
				for i := 0; i < len(userList); i++ {
					if userList[i] == req.UserId {
						userList = append(userList[:i], userList[i+1:]...)
					}
				}
				go addRelationToLike(strVideoId, userList)
			}
		}
		// 向消息队列发送消息
		rabbitmq.RmqUnLike.Publish(sb.String())
	}
	return resp, nil
}

// 将视频的点赞信息写入Redis
func addRelationToLike(key string, list []int64) {

	if _, err := redis.RdbLike.SAdd(redis.RedisCtx, key, util.RedisDefaultValue).Result(); err != nil {
		redis.RdbLike.Del(redis.RedisCtx, key)
		return
	}
	//给键值设置有效期，类似于gc机制
	if _, err := redis.RdbLike.Expire(redis.RedisCtx, key, time.Duration(util.OneDay)*time.Second).Result(); err != nil {
		redis.RdbLike.Del(redis.RedisCtx, key)
		return
	}
	for _, v := range list {
		if _, err := redis.RdbLike.SAdd(redis.RedisCtx, key, v).Result(); err != nil {
			redis.RdbLike.Del(redis.RedisCtx, key)
			return
		}
	}
}

func (l *LikeService) GetLikeList(ctx context.Context, req *service.LikeListRequest) (*service.LikeListResponse, error) {
	strUserId := util.LikeUserKey + strconv.FormatInt(req.UserId, 10)
	resp := new(service.LikeListResponse)
	resp.Code = util.Success

	return resp, nil
}
