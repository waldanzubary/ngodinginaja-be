package main

import (
	"log"
	"ngodinginaja-be/config"
	"ngodinginaja-be/models"
)

func main() {
	// 🔹 Hubungkan ke database dulu
	config.ConnectDB()
	db := config.DB

	log.Println("🧹 Dropping all tables...")
	err := db.Migrator().DropTable(
		&models.Submission{},
		&models.Lesson{},
		&models.Module{},
		&models.Course{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf("❌ Gagal drop tabel: %v", err)
	}

	log.Println("🚀 Migrating all tables...")
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

	log.Println("✅ Migrasi fresh berhasil!")
}
