package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"interaction/db"
	"interaction/pkg/util"
	"log"
	"strconv"
	"strings"
)

var RmqLike *LikeMQ
var RmqUnLike *LikeMQ

type LikeMQ struct {
	RabbitMQ
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

// InitLikeRabbitMQ 初始化rabbitMQ连接。
func InitLikeRabbitMQ() {
	RmqLike = NewLikeRabbitMQ("like_add")
	go RmqLike.Consumer()

	RmqUnLike = NewLikeRabbitMQ("like_del")
	go RmqUnLike.Consumer()
}

// NewLikeRabbitMQ 获取likeMQ的对应队列。
func NewLikeRabbitMQ(queueName string) *LikeMQ {
	likeMQ := &LikeMQ{
		RabbitMQ:  *Rmq,
		queueName: queueName,
	}
	cha, err := likeMQ.conn.Channel()
	likeMQ.channel = cha
	Rmq.failOnErr(err, "获取通道失败")
	return likeMQ
}

// Consumer like操作消费消息
func (l *LikeMQ) Consumer() {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := l.channel.QueueDeclare(
		l.queueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		panic(err)
	}

	//接收消息
	msgs, err1 := l.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err1 != nil {
		panic(err1)
	}

	forever := make(chan bool)
	//启用协程处理消息
	switch l.queueName {
	case "like_add":
		//点赞消费队列
		go l.consumerLikeAdd(msgs)
	case "like_del":
		//取消赞消费队列
		go l.consumerLikeDel(msgs)

	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

// Publish like操作的发布消息。
func (l *LikeMQ) Publish(message string) {

	_, err := l.channel.QueueDeclare(
		l.queueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		panic(err)
	}

	err1 := l.channel.Publish(
		l.exchange,
		l.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err1 != nil {
		panic(err)
	}

}

func (l *LikeMQ) consumerLikeAdd(messages <-chan amqp.Delivery) {
	for d := range messages {
		// 解析消息
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		userId, _ := strconv.ParseInt(params[0], 10, 64)
		videoId, _ := strconv.ParseInt(params[1], 10, 64)
		likeDao := db.TbLike{UserId: userId, VideoId: videoId, IsDel: util.ISLIKE}
		for i := 0; i < util.MaxAttempts; i++ {
			flag := true
			// 操作数据库
			info, err := likeDao.GetLikeInfo()
			if err != nil {
				// todo 打印日志
				flag = false
				continue
			}
			tx := db.Db.Begin()
			if info == nil {
				if err1 := likeDao.InsertLike(tx); err1 != nil {
					tx.Rollback()
					// todo 打印日志
					flag = false
					continue
				}
				exec := tx.Exec("Update video set favorite_count = favorite_count + 1 where id = ?", videoId)
				if exec.Error != nil {
					tx.Rollback()
					// todo 打印日志
					flag = false
					continue
				}
				tx.Commit()
			} else {
				if err1 := likeDao.UpdateLike(tx); err1 != nil {
					tx.Rollback()
					// todo 打印日志
					flag = false
					continue
				}
				exec := tx.Exec("Update video set favorite_count = favorite_count + 1 where id = ?", videoId)
				if exec.Error != nil {
					tx.Rollback()
					// todo 打印日志
					flag = false
					continue
				}
				tx.Commit()
			}
			if flag {
				break
			}
		}
	}
}

func (l *LikeMQ) consumerLikeDel(messages <-chan amqp.Delivery) {
	for d := range messages {
		// 解析消息
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		userId, _ := strconv.ParseInt(params[0], 10, 64)
		videoId, _ := strconv.ParseInt(params[1], 10, 64)
		likeDao := db.TbLike{UserId: userId, VideoId: videoId, IsDel: util.ISLIKE}
		for i := 0; i < util.MaxAttempts; i++ {
			flag := true
			info, err := likeDao.GetLikeInfo()
			if err != nil {
				// todo 打印日志
				flag = false
				continue
			}
			if info == nil {
				// todo 打印日志

				break
			}
			tx := db.Db.Begin()
			if err1 := likeDao.UpdateLike(tx); err1 != nil {
				tx.Rollback()
				// todo 打印日志
				flag = false
				continue
			}
			exec := tx.Exec("Update video set favorite_count = favorite_count - 1 where id = ?", videoId)
			if exec.Error != nil {
				tx.Rollback()
				// todo 打印日志
				flag = false
				continue
			}
			tx.Commit()
			if flag {
				break
			}
		}
	}
}
