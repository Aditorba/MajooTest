package repository

import (
	"gorm.io/gorm"
	"majooTest/domain"
)

type OutletRepository interface {
	SaveOutlet(data domain.Outlets) (domain.Outlets, error)
	GetAllData() ([]domain.Outlets, error)
	GetDetailOutlet(id int) (domain.Outlets, error)
	GetOutletByMerchantId(merchantId int) ([]domain.Outlets, error)
	DeleteOutlet(id int) (domain.Outlets, error)
}

type outletConnection struct {
	connection *gorm.DB
}

func NewOutletRepository(db *gorm.DB) OutletRepository {
	return &outletConnection{
		connection: db,
	}
}

func (db *outletConnection) SaveOutlet(data domain.Outlets) (domain.Outlets, error) {
	if err := db.connection.Save(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (db *outletConnection) GetAllData() ([]domain.Outlets, error) {
	var data []domain.Outlets
	if err := db.connection.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *outletConnection) GetDetailOutlet(id int) (domain.Outlets, error) {
	var data domain.Outlets
	if err := db.connection.Where("id=?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *outletConnection) GetOutletByMerchantId(merchantId int) ([]domain.Outlets, error) {
	var data []domain.Outlets
	if err := db.connection.Where("merchant_id=?", merchantId).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *outletConnection) DeleteOutlet(id int) (domain.Outlets, error) {
	var data domain.Outlets
	if err := db.connection.Where("id=?", id).Delete(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
