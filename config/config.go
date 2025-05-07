package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
}

var Cfg *Config

func Init() {
	// Load .env file jika ada
	if err := loadEnvFile(); err != nil {
		logrus.Warnf("Warning: File .env tidak ditemukan: %v", err)
	}

	viper.AddConfigPath(".")         

	viper.AutomaticEnv()


	Cfg = &Config{
		DB_NAME: getEnvOrConfig("DB_NAME", "task_tracker_db"),
		DB_USER: getEnvOrConfig("DB_USER", "task_tracker"),
		DB_PASS: getEnvOrConfig("DB_PASS", "task_tracker"),
		DB_HOST: getEnvOrConfig("DB_HOST", "localhost"),
		DB_PORT: getEnvOrConfig("DB_PORT", "5432"),
	}
}

// loadEnvFile mencoba memuat file .env dari beberapa lokasi
func loadEnvFile() error {
	envFiles := []string{".env", "../.env", "../../.env"}
	
	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			logrus.Infof("Berhasil memuat file .env dari: %s", file)
			return nil
		}
	}
	
	return godotenv.Load() // mencoba load dari lokasi default
}

// getEnvOrConfig mengambil nilai dari environment variable terlebih dahulu,
// jika tidak ada baru mengambil dari viper, dan terakhir menggunakan nilai default
func getEnvOrConfig(key string, defaultValue string) string {
	// Cek environment variable terlebih dahulu
	if value := os.Getenv(key); value != "" {
		return value
	}
	
	// Kemudian cek di viper
	if value := viper.GetString(key); value != "" {
		return value
	}
	
	// Jika tidak ada, gunakan nilai default
	return defaultValue
}