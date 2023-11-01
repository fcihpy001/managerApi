package service

import (
	"ManagerApi/model"
	"ManagerApi/utils"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db  *gorm.DB
	rdb *redis.Client
	err error
)

func createTable(db *gorm.DB) {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.ActiveCode{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.LoginRecord{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.Trade{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.Contract{}); err != nil {
		log.Printf("建表时出现异常", err)
	}

	if err := db.AutoMigrate(&model.Wallet{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.NFT{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.USDT{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.Chain{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.Reward{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.Rank{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.Deposit{}); err != nil {
		log.Printf("建表时出现异常", err)
	}
	if err := db.AutoMigrate(&model.Transaction{}); err != nil {
		log.Printf("建表时出现异常", err)
	}

}

func GetDB() *gorm.DB {
	if db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s",
			utils.Config.Datasource.UserName,
			utils.Config.Datasource.Password,
			utils.Config.Datasource.Host,
			utils.Config.Datasource.Database,
			utils.Config.Datasource.Charset, utils.Config.Datasource.Loc)
		DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: gormlogger.Default.LogMode(gormlogger.Info),
		})
		if err != nil {
			panic("failed to connect database")
		}
		db = DB
		logger.Info().Msg("数据库初始化成功...")
		createTable(db)
		log.Info().Msg("数据库建表成功...")
	}
	return db
}

//func initRedis(uri string) {
//	opts, err := redis.ParseURL(uri)
//	if err != nil {
//		panic(err)
//	}
//
//	rdb = redis.NewClient(opts)
//	if err = rdb.Ping(context.Background()).Err(); err != nil {
//		panic(err)
//	}
//	logger.Info().Msg("Redis 连接成功")
//
//}

var logger = log.With().Str("module", "services").Logger()
