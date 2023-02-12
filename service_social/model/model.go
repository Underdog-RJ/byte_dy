package model

import (
	"time"
)

type Message struct {
	ID         uint      `json:"id"`
	FromId     uint      `json:"from_id"`
	ToId       uint      `json:"to_id"`
	Content    string    `json:"content"`
	SendTime   time.Time `json:"send_time"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
}

type User struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	UserName       string    `json:"user_name"`
	PasswordDigest string    `json:"password_digest"`
}
type Follower struct {
	ID         uint      `json:"id"`
	FollowerId uint      `json:"follower_id"`
	FolloweeId uint      `json:"followee_id"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"crete_time"`
}

type UserList struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	FollowerCount uint   `json:"follower_count"`
	FollowCount   uint   `json:"follow_count"`
	IsFollow      bool   `json:"is_follow"`
}

type ListVO struct {
	StatusCode string     `json:"status_code"`
	StatusMsg  string     `json:"status_msg"`
	UserList   []UserList `json:"user_list"`
}

type MessageVo struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

type MessageListVO struct {
	StatusCode  string      `json:"status_code"`
	StatusMsg   string      `json:"status_msg"`
	MessageList []MessageVo `json:"message_list"`
}
