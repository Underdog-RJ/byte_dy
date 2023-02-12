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
	"sync"
	"time"
)

type LikeService struct{}

func NewLikeService() *LikeService {
	return new(LikeService)
}

func (l *LikeService) IsLike(ctx context.Context, req *service.IsLikeRequest) (*service.IsLikeResponse, error) {
	strUserId := util.LikeUserKey + strconv.FormatInt(req.UserId, 10)
	strVideoId := util.LikeVideoKey + strconv.FormatInt(req.VideoId, 10)
	resp := new(service.IsLikeResponse)
	resp.Code = util.Success

	if n, err := redis.RdbLike.Exists(ctx, strUserId).Result(); err != nil {
		resp.IsLike = false
		resp.Code = util.Error
		return resp, err
	} else {
		if n > 0 {
			if exist, err1 := redis.RdbLike.SIsMember(ctx, strUserId, req.VideoId).Result(); err1 != nil {
				resp.IsLike = false
				resp.Code = util.Error
				return resp, err
			} else {
				resp.IsLike = exist
				return resp, nil
			}
		}
	}
	if n, err := redis.RdbLike.Exists(ctx, strVideoId).Result(); err != nil {
		resp.IsLike = false
		resp.Code = util.Error
		return resp, err
	} else {
		if n > 0 {
			if exist, err1 := redis.RdbLike.SIsMember(ctx, strVideoId, req.UserId).Result(); err1 != nil {
				resp.IsLike = false
				resp.Code = util.Error
				return resp, err
			} else {
				resp.IsLike = exist
				return resp, nil
			}
		}
	}
	likeDao := db.TbLike{UserId: req.UserId, VideoId: req.VideoId}
	info, err := likeDao.GetLikeInfo()
	if err != nil {
		resp.IsLike = false
		resp.Code = util.Error
		return resp, err
	}
	resp.IsLike = info != nil

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
	strTarUserId := util.LikeUserKey + strconv.FormatInt(req.TargetUserId, 10)
	resp := new(service.LikeListResponse)
	resp.Code = util.Success

	if n, err := redis.RdbLike.Exists(ctx, strTarUserId).Result(); n > 0 {
		if err != nil {
			// todo 打印日志
			resp.Code = util.Error
			return resp, err
		}
		videoIdList, err1 := redis.RdbLike.SMembers(ctx, strTarUserId).Result()
		if err1 != nil {
			// todo 打印日志
			resp.Code = util.Error
			return resp, err1
		}
		// 通过数据库查询videoInfo
		var wg sync.WaitGroup
		i := len(videoIdList) - 1
		favoriteList := make([]*service.VideoInfo, i-1)
		if i == 0 {
			resp.VideoList = favoriteList
			return resp, nil
		}
		wg.Add(i)
		for index, videoId := range videoIdList {
			if videoId == util.RedisDefaultValue {
				continue
			}
			go addVideoInfoToList(videoId, favoriteList, req.UserId, &wg, index)
		}
		wg.Wait()
		resp.VideoList = favoriteList
		return resp, nil
	}

	return resp, nil
}

// addVideoInfoToList 添加视频信息到列表favoriteList中
func addVideoInfoToList(strVideoId string, favoriteList []*service.VideoInfo, userId int64, wg *sync.WaitGroup, index int) {
	defer wg.Done()
	videoId, _ := strconv.ParseInt(strVideoId, 10, 64)
	v := new(service.VideoInfo)
	row := db.Db.Raw("select v.id video_id, u.id user_id, u.user_name, u.follow_count, u.follower_count, v.play_url, v.cover_url ,v.favorite_count, v.comment_count, v.title from video v left join user u on u.id = v.user_id where v.id = ?", videoId).Row()
	err := row.Scan(&v.Id, &v.UserInfo.Id, v.UserInfo.Name, &v.UserInfo.FollowCount, &v.UserInfo.FollowerCount, &v.PlayUrl, &v.CoverUrl, &v.FavouriteCount, &v.CommentCount, &v.Title)
	if err != nil {
		// todo 打印日志
		return
	}
	req := service.IsLikeRequest{UserId: userId, VideoId: videoId}
	ctx := context.Background()
	likeService := LikeService{}
	resp, err := likeService.IsLike(ctx, &req)
	if err != nil {
		// todo 打印日志
		return
	}
	v.IsFavorite = resp.IsLike
	v.UserInfo.IsFollow = false
	var count int64
	db.Db.Raw("select count(1) from follower where followee_id = ? and follower_id = ?").Scan(&count)
	if count > 0 {
		v.UserInfo.IsFollow = true
	}
	favoriteList[index] = v
}
