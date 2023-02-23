package handle

import (
	"context"
	"errors"
	"interaction/db"
	"interaction/pkg/util"
	"interaction/service/service"
)

type CommentService struct{}

func NewCommentService() *CommentService {
	return new(CommentService)
}

func (c *CommentService) CommentAction(ctx context.Context, req *service.CommentRequest, resp *service.CommentResponse) error {
	var err error
	userId := req.UserId
	videoId := req.VideoId
	actionType := req.ActionType
	if err = db.FindUser(userId); err != nil {
		return err
	}
	if err = db.FindVideo(videoId); err != nil {
		return err
	}
	switch actionType {
	//发表评论
	case 1:
		commentText := req.CommentText
		comment := &db.Comment{
			UserId:      userId,
			VideoId:     videoId,
			CommentText: commentText,
			IsDel:       util.CommentIsNotDel,
		}
		comment, err = comment.InsertComment()
		if err != nil {
			return err
		}

		userInfo := db.UserInfo(userId)
		user := &service.User{
			Id:            userId,
			Name:          userInfo.UserName,
			FollowCount:   int64(userInfo.Follow_count),
			FollowerCount: int64(userInfo.Follower_count),
			IsFollow:      true,
		}
		respComment := &service.Comment{
			Id:         comment.Id,
			User:       user,
			Content:    comment.CommentText,
			CreateDate: comment.CreateTime.Format("01-02"),
		}
		resp.StatusCode = 0
		resp.StatusMsg = "评论发表成功"
		resp.Comment = respComment
	//删除评论
	case 2:
		commentId := req.CommentId
		comment := &db.Comment{
			Id: commentId,
		}
		err = comment.DeleteComment()
		if err != nil {
			return err
		}
		resp.StatusCode = 0
		resp.StatusMsg = "评论删除成功"
		resp.Comment = nil
	default:
		return errors.New("请求参数不合法")
	}
	return nil
}
