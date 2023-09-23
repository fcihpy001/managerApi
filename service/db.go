package service

import (
	"ManagerApi/model"
	"ManagerApi/utils"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

var (
	db  *gorm.DB
	rdb *redis.Client
	err error
)

func Init(ctx context.Context) {
	InitDB()
}

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.Config.Mysql.UserName,
		utils.Config.Mysql.Passwd,
		utils.Config.Mysql.Address,
		utils.Config.Mysql.Database)
	// fmt.Println(dsn)
	InitMysql(dsn)
	if os.Getenv("debug") == "1" {
		createTable()
	}

	//initRedis(model.Config.)
}

func InitMysql(dsn string) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	logger.Info().Msg("数据库初始化成功...")
}

//lint:ignore U1000 ignore unused lint
func createTable() {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		logger.Error().Stack().Err(err)
	}

}

func initRedis(uri string) {
	opts, err := redis.ParseURL(uri)
	if err != nil {
		panic(err)
	}

	rdb = redis.NewClient(opts)
	if err = rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	logger.Info().Msg("Redis 连接成功")

}

var logger = log.With().Str("module", "services").Logger()
