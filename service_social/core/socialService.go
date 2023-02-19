package core

import (
	"context"
	"errors"
	"log"
	"service_common/pkg/utils"
	"service_common/services"
	"service_social/model"
	"time"
)

func (*SocialService) RelationService(ctx context.Context, request *services.DouyinRelationActionRequest, response *services.DouyinRelationActionResponse) error {
	token := request.Token
	toUserID := request.ToUserId
	actionType := request.ActionType
	clim, _ := utils.ParseToken(token)
	userId := clim.Id
	response.StatusCode = 0
	response.StatusMsg = "success"
	follower := model.Follower{FollowerId: uint(toUserID), FolloweeId: userId, CreateTime: time.Now(), UpdateTime: time.Now()}
	if actionType == 1 {
		log.Println("关注操作")
		result := model.DB.Create(&follower)
		return result.Error
	} else {
		log.Println("取消关注关注操作")
		result := model.DB.Where("follower_id = ?", toUserID).Where("followee_id = ?", userId).Delete(&follower)
		return result.Error
	}

}
func BuildUser(result model.UserList) *services.User {
	userModel := services.User{
		Id:            int64(result.Id),
		Name:          result.Name,
		IsFollow:      result.IsFollow,
		FollowCount:   int64(result.FollowCount),
		FollowerCount: int64(result.FollowerCount),
	}
	return &userModel
}
func (*SocialService) RelationFollowList(ctx context.Context, request *services.DouyinRelationFollowListRequest, response *services.DouyinRelationFollowListResponse) error {
	token := request.Token
	myUserId := request.UserId
	log.Println("开始返回关注人数")
	followers := []model.Follower{}
	user := model.User{}
	results := []model.UserList{}
	clim, _ := utils.ParseToken(token)
	//myUserId, _ := strconv.Atoi(userId)
	log.Println("token{}", clim.Id)
	response.StatusCode = 0
	response.StatusMsg = "succeess"
	mr := model.DB.Where("followee_id = ?", myUserId).Find(&followers)
	if mr.RowsAffected == 0 {
		return errors.New("查询结果为nil")
	} else {
		for i := 0; i < len(followers); i++ {
			result := model.UserList{}
			userID := followers[i].FollowerId
			result.Id = userID
			user = model.User{}
			model.DB.Where("id = ?", userID).Find(&user)
			log.Println("user:", user)
			result.Name = user.UserName

			//计算粉丝数和关注数
			Relation(userID, uint(myUserId), &result)

			results = append(results, result)
		}

		for i := 0; i < len(results); i++ {
			response.UserList = append(response.UserList, BuildUser(results[i]))
		}

		return nil
	}
}

func (*SocialService) RelationFollowerList(ctx context.Context, in *services.DouyinRelationFollowerListRequest, out *services.DouyinRelationFollowerListResponse) error {
	token := in.Token
	myUserId := in.UserId
	log.Println("开始返回粉丝数量")
	followers := []model.Follower{}
	user := model.User{}

	results := []model.UserList{}
	clim, _ := utils.ParseToken(token)
	//myUserId, _ := strconv.Atoi(userId)
	log.Println("token:", clim.Id)
	log.Println("myUserId:", myUserId)
	out.StatusMsg = "success"
	out.StatusCode = 0
	mr := model.DB.Where("follower_id = ?", myUserId).Find(&followers)
	if mr.RowsAffected == 0 {
		return errors.New("查询结果为nil")
	} else {
		for i := 0; i < len(followers); i++ {
			result := model.UserList{}
			userID := followers[i].FolloweeId
			log.Println("followers:", followers)
			log.Println("id:", followers[i].ID)
			result.Id = userID
			log.Println("userID:", userID)

			user = model.User{}
			model.DB.Where("id = ?", userID).Find(&user)
			log.Println("user:", user)
			result.Name = user.UserName

			//计算粉丝和关注人数
			Relation(userID, uint(myUserId), &result)

			results = append(results, result)
		}
		for i := 0; i < len(results); i++ {
			out.UserList = append(out.UserList, BuildUser(results[i]))
		}

		return nil
	}
}

// 计算粉丝数和关注数
func Relation(userID uint, myUserId uint, result *model.UserList) {
	followerss := []model.Follower{}
	r := model.DB.Where("follower_id = ?", userID).Find(&followerss)
	followerCount := r.RowsAffected
	result.FollowerCount = uint(followerCount)
	fR := model.DB.Where("followee_id = ?", userID).Find(&followerss)
	//关注的人
	followCount := fR.RowsAffected
	log.Println("关注了：", fR.RowsAffected)
	result.FollowCount = uint(followCount)
	//我是否关注
	fBool := model.DB.Where("follower_id = ?", userID).Where("followee_id", myUserId).Find(&followerss)
	if fBool.RowsAffected == 1 {
		result.IsFollow = true
	} else if fBool.RowsAffected == 0 {
		result.IsFollow = false
	}
}

func (*SocialService) RelationFriendList(ctx context.Context, in *services.DouyinRelationFriendListRequest, out *services.DouyinRelationFriendListResponse) error {
	token := in.Token
	myUserId := in.UserId
	log.Println("开始返回好友数量")
	followers := []model.Follower{}
	user := model.User{}
	followerss := []model.Follower{}
	results := []model.UserList{}
	followersss := []model.Follower{}
	clim, _ := utils.ParseToken(token)
	out.StatusCode = 0
	out.StatusMsg = "success"
	//myUserId, _ := strconv.Atoi(userId)
	log.Println("token:", clim.Id)
	log.Println("myUserId:", myUserId)
	mr := model.DB.Where("follower_id = ?", myUserId).Find(&followers)
	if mr.RowsAffected == 0 {
		return errors.New("查询结果为nil")
	} else {
		for i := 0; i < len(followers); i++ {

			userID := followers[i].FolloweeId
			log.Println("userID:", userID)
			model.DB.Where("follower_id = ?", userID).Find(&followerss)
			for i := 0; i < len(followerss); i++ {
				followeeId := followerss[i].FolloweeId
				if followeeId == uint(myUserId) {
					//好友数量
					followersss = append(followersss, followerss[i])
				}
			}
		}
		log.Println("好友：", followersss)
		for i := 0; i < len(followersss); i++ {
			userID := followersss[i].FollowerId
			result := model.UserList{}
			result.Id = userID
			user = model.User{}
			model.DB.Where("id = ?", userID).Find(&user)
			result.Name = user.UserName

			//计算粉丝数和关注数量
			Relation(userID, uint(myUserId), &result)
			log.Println("result:", result)
			results = append(results, result)
		}

		//
		//for i := 0; i < len(results); i++ {
		//
		//	out.UserList = append(out.UserList, BuildUser(results[i]))
		//}

		return nil
	}
}

// 发送消息
func (*SocialService) MessageAction(ctx context.Context, in *services.DouyinRelationActionRequestContent, out *services.DouyinRelationActionResponse) error {
	token := in.Token
	toUserID := in.ToUserId
	actionType := in.ActionType
	content := in.Content
	clim, _ := utils.ParseToken(token)
	myUserId := clim.Id
	out.StatusMsg = "success"
	out.StatusCode = 0

	//toUserID, _ := strconv.Atoi(toUserId)
	if actionType == 1 {
		message := model.Message{FromId: myUserId, ToId: uint(toUserID), SendTime: time.Now(), Content: content, UpdateTime: time.Now(), CreateTime: time.Now()}
		model.DB.Create(&message)
		return nil
	} else {
		err := errors.New("发送失败！")
		panic(err)
	}
}

// 聊天记录
func (*SocialService) MessageChat(ctx context.Context, in *services.DouyinMessageChatRequest, out *services.DouyinMessageChatResponse) error {
	token := in.Token
	toUserId := in.ToUserId
	out.StatusCode = 0
	out.StatusMsg = "success"
	log.Println("获取聊天记录")
	clim, _ := utils.ParseToken(token)
	myUserId := clim.Id
	messages := []model.Message{}
	model.DB.Where("from_id = ?", myUserId).Where("to_id = ?", toUserId).Find(&messages)
	for i := 0; i < len(messages); i++ {
		//result := model.MessageVo{}
		result := services.Message{}

		message := messages[i]
		result.Id = int64(message.ID)
		result.Content = message.Content
		result.ToUserId = toUserId
		result.FromUserId = int64(myUserId)
		result.CreateTime = (message.CreateTime).Format("2006-01-02 15:04:05")

		out.MessageList = append(out.MessageList, &result)

	}

	return nil
}
