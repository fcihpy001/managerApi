package utils

import (
	"ManagerApi/model"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	Config model.Config
)

func InitConfig() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Fatalf("无法加载 .env 文件: %v", err)
	}
	log.Println("env文件加载成功")

	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		return
	}
	log.Println("配置文件读取成功")
}

func RandString(length int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, length)

	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func RandStringAndNumber(length int) string {
	var letters = []byte("ABCDEF1234567890")
	result := make([]byte, length)

	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return strings.ToUpper(string(result))
}

func HashStr(str1 string, str2 string, str3 string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str1))
	hasher.Write([]byte(str2))
	hasher.Write([]byte(str3))
	hashBytes := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hashBytes)
	return hashHex
}
