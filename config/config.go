package config

import (
	"blog-backend/models"
	"os"

	//"blog-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func init() {
	// 尝试加载 .env 文件（开发环境）
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env")
	}
}

func GetSecretKey() string {
	// 1. 优先从环境变量读取
	if key := os.Getenv("AUTH_SECRET_KEY"); key != "" {
		return key
	}

	// 2. 回退到配置文件（不推荐生产环境使用）
	log.Fatal("AUTH_SECRET_KEY must be set in environment variables")
	return ""
}

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	log.Println("Database connection successfully opened")

	// 自动迁移模型
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("Failed to migrate database!")
	}
}
