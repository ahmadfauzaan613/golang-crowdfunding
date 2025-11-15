package databases

import (
	"crowdfunding/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBconnection(cfg *config.Config) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Failed GORM Open:", err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get sqlDB:", err)
		return nil
	}

	if err := sqlDB.Ping(); err != nil {
		log.Println("Ping failed:", err)
		return nil
	}

	return db
}
