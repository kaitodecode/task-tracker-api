package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDatabase() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		Cfg.DB_HOST,
		Cfg.DB_USER,
		Cfg.DB_PASS,
		Cfg.DB_NAME,
		Cfg.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorf("Gagal koneksi ke database: %v", err)
		panic(err)
	}
	
	// Set connection pool
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorf("Gagal mengatur connection pool: %v", err)
		panic(err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	logrus.Info("Koneksi database berhasil dibuat")
	return db
}
