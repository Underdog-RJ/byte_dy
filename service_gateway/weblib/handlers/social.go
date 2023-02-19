package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"service_common/services"
	"strconv"
)

var Relation = new(relation)

type relation struct{}

// 关注操作
func (r *relation) RelationAction(ginCtx *gin.Context) {
	token := ginCtx.Query("token")
	toUserId := ginCtx.Query("to_user_id")
	actionType := ginCtx.Query("action_type")
	var socialReq services.DouyinRelationActionRequest
	socialReq.Token = token
	toUser, _ := strconv.ParseInt(toUserId, 10, 64)
	socialReq.ToUserId = toUser
	act, _ := strconv.ParseInt(actionType, 10, 32)

	socialReq.ToUserId = toUser
	socialReq.ActionType = int32(act)
	// 从gin.Key中取出服务实例
	socialService := ginCtx.Keys["socialService"].(services.SocialService)

	socialResp, _ := socialService.RelationService(context.Background(), &socialReq)

	ginCtx.JSON(200, gin.H{"status_code": socialResp.GetStatusCode(), "status_msg": socialResp.GetStatusMsg()})
}

// 关注列表
func (r *relation) RelationFollowList(ginCtx *gin.Context) {
	userId := ginCtx.Query("user_id")
	token := ginCtx.Query("token")
	var socialReq services.DouyinRelationFollowListRequest
	userI, _ := strconv.ParseInt(userId, 10, 64)
	socialReq.UserId = userI
	socialReq.Token = token
	socialService := ginCtx.Keys["socialService"].(services.SocialService)
	socialResp, _ := socialService.RelationFollowList(context.Background(), &socialReq)

	ginCtx.JSON(200, gin.H{"status_code": socialResp.GetStatusCode(), "status_msg": socialResp.GetStatusMsg(), "user_list": socialResp.UserList})
}

// 粉丝列表
func (r *relation) RelationFollowerList(ginCtx *gin.Context) {
	userId := ginCtx.Query("user_id")
	token := ginCtx.Query("token")

	var socialReq services.DouyinRelationFollowerListRequest
	userI, _ := strconv.ParseInt(userId, 10, 64)
	socialReq.UserId = userI
	socialReq.Token = token
	socialService := ginCtx.Keys["socialService"].(services.SocialService)
	socialResp, _ := socialService.RelationFollowerList(context.Background(), &socialReq)

	ginCtx.JSON(200, gin.H{"status_code": socialResp.GetStatusCode(), "status_msg": socialResp.GetStatusMsg(), "user_list": socialResp.UserList})
}

// 好友列表
func (r *relation) RelationFriendList(ginCtx *gin.Context) {
	userId := ginCtx.Query("user_id")
	token := ginCtx.Query("token")
	var socialReq services.DouyinRelationFriendListRequest
	userI, _ := strconv.ParseInt(userId, 10, 64)
	socialReq.UserId = userI
	socialReq.Token = token
	socialService := ginCtx.Keys["socialService"].(services.SocialService)
	socialResp, _ := socialService.RelationFriendList(context.Background(), &socialReq)
	ginCtx.JSON(200, gin.H{"status_code": socialResp.GetStatusCode(), "status_msg": socialResp.GetStatusMsg()})
}

// 发送消息
func (r *relation) MessageAction(ginCtx *gin.Context) {
	token := ginCtx.Query("token")
	toUserId := ginCtx.Query("to_user_id")
	actionType := ginCtx.Query("action_type")
	content := ginCtx.Query("content")
	var socialReq services.DouyinRelationActionRequestContent
	socialReq.Token = token
	toUser, _ := strconv.ParseInt(toUserId, 10, 64)
	socialReq.ToUserId = toUser
	act, _ := strconv.ParseInt(actionType, 10, 64)
	socialReq.ActionType = int32(act)
	socialReq.Content = content
	socialService := ginCtx.Keys["socialService"].(services.SocialService)

	socialResp, err := socialService.MessageAction(context.Background(), &socialReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status_code": socialResp.StatusCode, "status_msg": err.Error()})
	} else {
		ginCtx.JSON(200, gin.H{"status_code": socialResp.StatusCode, "status_msg": "发送成功"})
	}
}

// 聊天记录
func (r *relation) MessageChat(ginCtx *gin.Context) {
	token := ginCtx.Query("token")
	toUserId := ginCtx.Query("to_user_id")
	socialService := ginCtx.Keys["socialService"].(services.SocialService)
	var socialReq services.DouyinMessageChatRequest
	socialReq.Token = token
	toUser, _ := strconv.ParseInt(toUserId, 10, 64)
	socialReq.ToUserId = toUser
	socialResp, _ := socialService.MessageChat(context.Background(), &socialReq)

	ginCtx.JSON(200, gin.H{"status_code": socialResp.StatusCode, "status_msg": "发送成功", "message_list": socialResp.MessageList})
}
