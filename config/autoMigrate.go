package config

import (
	"gorm.io/gorm"
	"majooTest/domain"
)

func Migration(db *gorm.DB) {

	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Merchants{})
	db.AutoMigrate(&domain.Outlets{})
	db.AutoMigrate(&domain.Transactions{})
}
