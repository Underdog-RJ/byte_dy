package conf

import (
	"fmt"
	"github.com/go-ini/ini"
	"service_video_processor/model"
	"strings"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string
)

type MinioStruct struct {
	MinioUrl       string
	MinioPort      string
	MinioAccessKey string
	MinioSecretKey string
}

// Init 初始化配置项
func Init() {
	// 连接数据库
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}

	// 初始化mysql
	LoadMysqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)

	// 连接RabbitMQ
	LoadRabbitMQ(file)
	pathRabbitMQ := strings.Join([]string{RabbitMQ, "://", RabbitMQUser, ":", RabbitMQPassWord, "@", RabbitMQHost, ":", RabbitMQPort, "/"}, "")
	model.RabbitMQ(pathRabbitMQ)

	// 初始化Minio
	client := LoadMinioData(file)
	model.Minio(client)
}

func LoadRabbitMQ(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadMinioData(file *ini.File) map[string]string {
	minioClient := make(map[string]string)
	minioClient["MinioUrl"] = file.Section("Minio").Key("MinioUrl").String()
	minioClient["MinioPort"] = file.Section("Minio").Key("MinioPort").String()
	minioClient["MinioAccessKey"] = file.Section("Minio").Key("MinioAccessKey").String()
	minioClient["MinioSecretKey"] = file.Section("Minio").Key("MinioSecretKey").String()
	return minioClient
}
