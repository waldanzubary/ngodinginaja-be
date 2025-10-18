package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ngodinginaja-be/models"
)

var DB *gorm.DB

func ConnectDB() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Tidak menemukan file .env, pastikan file-nya ada di root project")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal konek database: %v", err)
	}

	DB = db

	err = db.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Module{},
		&models.Lesson{},
		&models.Submission{},
	)
	if err != nil {
		log.Fatalf("❌ Gagal migrasi: %v", err)
	}

	log.Println("✅ Database terkoneksi & migrasi berhasil!")
}
