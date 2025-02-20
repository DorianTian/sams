package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitPgDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=tianqiyin dbname=sams password=<tqy> sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}
