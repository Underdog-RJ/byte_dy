package util

const (
	// 点赞操作
	ISLIKE = 0
	// 取消点赞操作
	UNLIKE = 1

	// 评论未删除
	CommentIsNotDel = 0
	// 评论已删除
	CommentIsDel = 1

	// Redis 存储用户喜欢列表的key的前缀
	LikeUserKey = "LikeUser:"
	// Redis 存储视频点赞列表的key的前缀
	LikeVideoKey = "LikeVideo:"
	// Redis set存储默认值
	RedisDefaultValue = -1

	// 返回错误状态码
	Error = 500
	// 返回正确状态码
	Success = 200

	// 一天的秒数
	OneDay = 24 * 60 * 60

	// 消息队列最大尝试次数
	MaxAttempts = 5
)
