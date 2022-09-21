package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func ConnectionDB(configData *ConfigData) *gorm.DB {

	dsn := configData.Database.DBName
	fmt.Println("Database name : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connection failed", err)
	} else {
		log.Println("connection db success")
	}

	Migration(db)
	return db
}
