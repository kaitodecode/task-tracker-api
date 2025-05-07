package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	logrus.Info("Proses migrasi dimulai")
	err := db.AutoMigrate(

	)
	if err != nil {
		logrus.Errorf("Error pada proses migrasi %v", err)
	}
	logrus.Info("Proses migrasi sukses")
}