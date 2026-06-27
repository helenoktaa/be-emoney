package database

import (
	"fmt"
	"log"
	"time"

	"emoney-2fa/config"
	"emoney-2fa/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySQL(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FJakarta",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	loc, _ := time.LoadLocation("Asia/Jakarta")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().In(loc)
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.OTP{},
		&models.Account{},
		&models.Transaction{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("MySQL connected and migrated")
	return db
}