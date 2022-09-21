package repository

import (
	"gorm.io/gorm"
	"majooTest/domain"
)

type MerchantRepository interface {
	SaveMerchant(data domain.Merchants) (domain.Merchants, error)
	GetAllData() ([]domain.Merchants, error)
	GetDetailMerchant(id int) (domain.Merchants, error)
	GetDetailMerchantByUserId(userId int) (domain.Merchants, error)
	DeleteMerchant(id int) (domain.Merchants, error)
}

type merchantConnection struct {
	connection *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantConnection{
		connection: db,
	}
}

func (db *merchantConnection) SaveMerchant(data domain.Merchants) (domain.Merchants, error) {
	if err := db.connection.Save(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (db *merchantConnection) GetAllData() ([]domain.Merchants, error) {
	var data []domain.Merchants
	if err := db.connection.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *merchantConnection) GetDetailMerchant(id int) (domain.Merchants, error) {
	var data domain.Merchants
	if err := db.connection.Where("id=?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *merchantConnection) GetDetailMerchantByUserId(userId int) (domain.Merchants, error) {
	var data domain.Merchants
	if err := db.connection.Where("user_id=?", userId).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *merchantConnection) DeleteMerchant(id int) (domain.Merchants, error) {
	var data domain.Merchants
	if err := db.connection.Where("id=?", id).Delete(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
