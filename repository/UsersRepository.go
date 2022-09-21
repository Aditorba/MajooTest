package repository

import (
	"gorm.io/gorm"
	"majooTest/domain"
)

type UsersRepository interface {
	SaveUser(data domain.Users) (domain.Users, error)
	GetAllData() ([]domain.Users, error)
	GetDetailUser(id int) (domain.Users, error)
	GetDetailByUsername(username string) (domain.Users, error)
	DeleteUser(id int) (domain.Users, error)
}

type usersConnection struct {
	connection *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersConnection{
		connection: db,
	}
}

func (db *usersConnection) SaveUser(data domain.Users) (domain.Users, error) {
	if err := db.connection.Save(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (db *usersConnection) GetAllData() ([]domain.Users, error) {
	var data []domain.Users
	if err := db.connection.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *usersConnection) GetDetailUser(id int) (domain.Users, error) {
	var data domain.Users
	if err := db.connection.Where("id=?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *usersConnection) GetDetailByUsername(username string) (domain.Users, error) {
	var data domain.Users
	if err := db.connection.Where("user_name=?", username).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *usersConnection) DeleteUser(id int) (domain.Users, error) {
	var data domain.Users
	if err := db.connection.Where("id=?", id).Delete(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
