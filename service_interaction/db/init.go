package db

import (
	"strings"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var Db *gorm.DB

func InitDB() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	addr := viper.GetString("mysql.addr")
	port := viper.GetString("mysql.port")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")

	dsn := strings.Join([]string{username, ":", password, "@tcp(", addr, ":", port, ")/", database, "?charset=utf8&parseTime=true&loc=Local"}, "")

	Db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      gormlogrus,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := Db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

	//Db.Set(`gorm:table_options`, "charset=utf8mb4").
	//	AutoMigrate(&Comment{})
}
