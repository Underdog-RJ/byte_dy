package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	mqurl string
}

const MQURL = "amqp://guest:Zhangzhengxu123.@127.0.0.1:5672/"

var Rmq *RabbitMQ

// InitRabbitMQ 初始化RabbitMQ的连接和通道。
func InitRabbitMQ() {

	Rmq = &RabbitMQ{
		mqurl: MQURL,
	}
	dial, err := amqp.Dial(Rmq.mqurl)
	Rmq.failOnErr(err, "创建连接失败")
	Rmq.conn = dial

}

// 连接出错时，输出错误信息。
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf("%s:%s\n", err, message))
	}
}

// 关闭mq通道和mq的连接。
func (r *RabbitMQ) destroy() {
	r.conn.Close()
}
