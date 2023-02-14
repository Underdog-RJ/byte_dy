package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service_common/pkg/logging"
	"service_common/pkg/utils"
	"service_common/services"
	"strconv"
)

type commonResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type likeListResponse struct {
	commonResponse
	vidoeList []videoinfo `json:"vidoe_list"`
}

func LikeAction(ginCtx *gin.Context) {
	token := ginCtx.Query("token")
	strVideoId := ginCtx.Query("video_id")
	strActionType := ginCtx.Query("action_type")
	if token == "" {
		log.Println("权限不足")
		ginCtx.JSON(http.StatusOK, commonResponse{
			StatusCode: 1,
			StatusMsg:  "权限不足",
		})
	}

	parseToken, _ := utils.ParseToken(token)
	userId := parseToken.Id
	videoId, _ := strconv.ParseInt(strVideoId, 10, 64)
	actionType, _ := strconv.Atoi(strActionType)

	req := services.LikeActionRequest{
		UserId:     int64(userId),
		VideoId:    videoId,
		ActionType: uint32(actionType),
	}
	likeService := ginCtx.Keys["likeService"].(services.LikeService)
	resp, err := likeService.LikeAction(ginCtx, &req)
	if err != nil {
		err = errors.New("likeService--" + err.Error())
		logging.Info(err)
		ginCtx.JSON(http.StatusOK, commonResponse{
			StatusCode: int(resp.Code),
			StatusMsg:  err.Error(),
		})
	}
	ginCtx.JSON(http.StatusOK, commonResponse{
		StatusCode: int(resp.Code),
		StatusMsg:  "ok",
	})
}

func LikeList(ginCtx *gin.Context) {
	token := ginCtx.Query("token")
	strTarUserId := ginCtx.Query("user_id")
	if token == "" {
		log.Println("权限不足")
		ginCtx.JSON(http.StatusOK, commonResponse{
			StatusCode: 1,
			StatusMsg:  "权限不足",
		})
	}
	parseToken, _ := utils.ParseToken(token)
	userId := parseToken.Id
	tarUserId, _ := strconv.ParseInt(strTarUserId, 10, 64)

	req := services.LikeListRequest{
		UserId:       int64(userId),
		TargetUserId: tarUserId,
	}

	likeService := ginCtx.Keys["likeService"].(services.LikeService)
	resp, err := likeService.GetLikeList(ginCtx, &req)
	if err != nil {
		err = errors.New("likeService--" + err.Error())
		logging.Info(err)
		ginCtx.JSON(http.StatusOK, commonResponse{
			StatusCode: int(resp.Code),
			StatusMsg:  err.Error(),
		})
	}

	ginCtx.JSON(http.StatusOK, likeListResponse{
		commonResponse: commonResponse{
			StatusCode: int(resp.Code),
			StatusMsg:  "ok",
		},
		vidoeList: buildVideo(resp.VideoList),
	})

}

func buildVideo(videoList []*services.VideoInfo) []videoinfo {
	list := make([]videoinfo, len(videoList))
	for i, v := range videoList {
		tmp := videoinfo{
			ID:            int(v.Id),
			Authorinfo:    toAuthorinfo(v.UserInfo),
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: int(v.FavouriteCount),
			CommentCount:  int(v.CommentCount),
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		}
		list[i] = tmp
	}
	return list
}

func toAuthorinfo(user *services.UserInfo) authorinfo {
	return authorinfo{
		Id:             int(user.Id),
		User_name:      user.Name,
		Follow_count:   int(user.FollowCount),
		Follower_count: int(user.FollowerCount),
		Is_follow:      user.IsFollow,
	}
}
