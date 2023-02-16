package handlers

import (
	"errors"
	"net/http"
	"service_common/pkg/logging"

	"github.com/gin-gonic/gin"
)

// 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

func PanicIfVideoError(err error) {
	if err != nil {
		err = errors.New("videoService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

func PanicIfCommentError(err error) {
	if err != nil {
		err = errors.New("commentService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

func UserPanicHandler(ginCtx *gin.Context) {
	if err := recover(); err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"status_code": 400,
			"status_msg":  err.(error).Error(),
		})
	}
}
