package core

import (
	"context"
	"errors"
	"user/model"
	"user/services"

	"github.com/jinzhu/gorm"
)

func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserResponse) error {
	var user model.User
	resp.Code = 0
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("用户不存在")
		}
		return err
	}
	if user.CheckPassword(req.Password) == false {
		return errors.New("密码错误")
	}
	resp.ID = uint32(user.ID)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserResponse) error {
	count := 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	user := model.User{
		UserName: req.UserName,
	}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	resp.ID = uint32(user.ID)
	resp.Code = 0
	return nil
}
