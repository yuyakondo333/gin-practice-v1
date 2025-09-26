package infra

import (
	"fmt"

	"github.com/you/myapp-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dsn := config.AppConfig.Database.GetDSN()

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		DB = nil
		return
	}
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() error {
	if DB == nil {
		return nil
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("データベース接続の取得に失敗しました: %v", err)
	}
	return sqlDB.Close()
}
