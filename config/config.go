package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GalaSetup() (*gorm.DB, error) {
	// GalaSetup : Set Connection to MySQL !
	strict := "root:root@tcp(localhost:3306)/sample_golang_apps"

	database, err := gorm.Open(mysql.Open(strict), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fail: %v", err)
		return nil, err
	}
	return database, nil
}
