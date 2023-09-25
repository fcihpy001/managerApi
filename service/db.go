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
)

var (
	DB  *gorm.DB
	rdb *redis.Client
	err error
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s",
		utils.Config.Datasource.UserName,
		utils.Config.Datasource.Password,
		utils.Config.Datasource.Host,
		utils.Config.Datasource.Database,
		utils.Config.Datasource.Charset, utils.Config.Datasource.Loc)
	log.Printf(fmt.Sprintf("dsn:", dsn))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	logger.Info().Msg("数据库初始化成功...")
	createTable()
	log.Info().Msg("数据库建表成功...")
}

func createTable() {
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := DB.AutoMigrate(&model.ActiveCode{}); err != nil {
		log.Printf("建表时出现异常", err)
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
