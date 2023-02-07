package handlers

import (
	"errors"
	"service_common/pkg/logging"
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
