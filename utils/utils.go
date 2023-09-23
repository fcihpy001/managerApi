package utils

import (
	"ManagerApi/model"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	Config model.Config
)

func InitConfig() {
	// 读取配置文件
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 解析 YAML 配置文件
	if err := yaml.Unmarshal(data, &model.Config{}); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
	log.Println("配置文件加载成功...")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if os.Getenv("debug") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Println("日志模块初始化成功...")
}
