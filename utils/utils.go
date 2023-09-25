package utils

import (
	"ManagerApi/model"
	"fmt"
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
	var letters = []byte("abcdefghjkmnpqrstuvwxyz123456789")
	result := make([]byte, length)

	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return strings.ToUpper(string(result))
}
