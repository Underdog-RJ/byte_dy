package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"service_video/model"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

type MinioStruct struct {
	MinioUrl       string
	MinioPort      string
	MinioAccessKey string
	MinioSecretKey string
}

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	//LoadMysqlData(file)
	//path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	//fmt.Println(path)
	//model.Database(path)
	client := LoadMinioData(file)
	model.Minio(client)
}

func LoadMinioData(file *ini.File) map[string]string {
	minioClient := make(map[string]string)
	minioClient["MinioUrl"] = file.Section("Minio").Key("MinioUrl").String()
	minioClient["MinioPort"] = file.Section("Minio").Key("MinioPort").String()
	minioClient["MinioAccessKey"] = file.Section("Minio").Key("MinioAccessKey").String()
	minioClient["MinioSecretKey"] = file.Section("Minio").Key("MinioSecretKey").String()
	return minioClient
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
