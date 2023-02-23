package handlers

import (
	"context"
	"net/http"
	"service_common/services/service"

	"github.com/gin-gonic/gin"
)

func CommentAction(ginCtx *gin.Context) {
	var commentReq service.CommentRequest
	PanicIfCommentError(ginCtx.Bind(&commentReq))
	userId, _ := ginCtx.Get("parseUserId")
	commentReq.UserId = userId.(int64)
	commentService := ginCtx.Keys["commentService"].(service.CommentService)
	commentActionResp, err := commentService.CommentAction(context.Background(), &commentReq)
	PanicIfCommentError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": commentActionResp.StatusCode,
		"status_msg":  commentActionResp.StatusMsg,
		"comment":     commentActionResp.Comment,
	})
}
