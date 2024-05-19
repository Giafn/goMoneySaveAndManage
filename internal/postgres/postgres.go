package postgres

import (
	"fmt"
	"log"

	"github.com/Giafn/goMoneySaveAndManage/configs"
	"github.com/Giafn/goMoneySaveAndManage/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		configs.AppConfig.DBHost,
		configs.AppConfig.DBUser,
		configs.AppConfig.DBPassword,
		configs.AppConfig.DBName,
		configs.AppConfig.DBPort,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
