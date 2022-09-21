package repository

import (
	"fmt"
	"gorm.io/gorm"
	"majooTest/domain"
	"majooTest/log"
	"math"
)

type TransactionRepository interface {
	SaveTransaction(data domain.Transactions) (domain.Transactions, error)
	GetAllData() ([]domain.Transactions, error)
	FindAllMerchantTransactionWithFilterPage(
		userId int, merchantId int64,
		outletId int64, startDate string,
		endDate string, page int64, limit int64) (resultData []domain.DataReportTemplate, totalData int64, currentPage int64, lastPage float64, err error)
	FindAllOutletTransactionWithFilterPage(
		userId int, merchantId int64,
		outletId int64, startDate string,
		endDate string, page int64, limit int64) (resultData []domain.DataReportTemplate, totalData int64, currentPage int64, lastPage float64, err error)
}

type transactionConnection struct {
	connection *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionConnection{
		connection: db,
	}
}

func (db *transactionConnection) SaveTransaction(data domain.Transactions) (domain.Transactions, error) {
	if err := db.connection.Save(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (db *transactionConnection) GetAllData() ([]domain.Transactions, error) {
	var data []domain.Transactions
	if err := db.connection.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *transactionConnection) FindAllMerchantTransactionWithFilterPage(
	userId int, merchantId int64,
	outletId int64, startDate string,
	endDate string, page int64, limit int64) (resultData []domain.DataReportTemplate, totalData int64, currentPage int64,
	lastPage float64, err error) {

	var dataList []domain.DataReportTemplate
	var total int64

	query := db.connection.Table("transactions as t").
		Select("t.id, t.merchant_id, m.merchant_name, t.outlet_id, o.outlet_name, "+
			"m.user_id, u.name, sum(t.bill_total) omzet").
		Joins("LEFT JOIN merchants as m ON t.merchant_id = m.id").
		Joins("LEFT JOIN outlets as o ON t.outlet_id = o.id").
		Joins("LEFT JOIN users as u ON m.user_id = u.id").
		Where("u.id =?", userId)

	if merchantId != 0 {
		query.Where("t.merchant_id = ?", merchantId)

		if outletId != 0 {
			query.Where("t.outlet_id = ?", outletId)
		}
	} else {
		if outletId != 0 {
			query.Where("t.outlet_id = ?", outletId)
		}
	}

	if startDate != "" {
		query.Where("t.created_at > ?", startDate)

		if endDate != "" {
			query.Where("t.created_at < ?", endDate)
		}
	} else {
		if endDate != "" {
			query.Where("t.created_at < ?", endDate)
		}
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 5
	}

	query.Limit(int(limit))

	offset := (page - 1) * limit
	query.Offset(int(offset))

	query.Group("t.merchant_id")
	query.Count(&total)

	lastPage = math.Ceil(float64(total / limit))
	if err := query.Scan(&dataList).Error; err != nil {
		return dataList, total, page, lastPage, err
	}

	log.Info("result data list :", dataList)
	fmt.Println("datalist : ", dataList)
	return dataList, total, page, lastPage, err
}

func (db *transactionConnection) FindAllOutletTransactionWithFilterPage(
	userId int, merchantId int64,
	outletId int64, startDate string,
	endDate string, page int64, limit int64) (resultData []domain.DataReportTemplate, totalData int64, currentPage int64,
	lastPage float64, err error) {

	var dataList []domain.DataReportTemplate
	var total int64

	query := db.connection.Table("transactions as t").
		Select("t.id, t.merchant_id, m.merchant_name, t.outlet_id, o.outlet_name, "+
			"m.user_id, u.name, sum(t.bill_total) omzet").
		Joins("LEFT JOIN merchants as m ON t.merchant_id = m.id").
		Joins("LEFT JOIN outlets as o ON t.outlet_id = o.id").
		Joins("LEFT JOIN users as u ON m.user_id = u.id").
		Where("u.id =?", userId)

	if merchantId != 0 {
		query.Where("t.merchant_id = ?", merchantId)

		if outletId != 0 {
			query.Where("t.outlet_id = ?", outletId)
		}
	} else {
		if outletId != 0 {
			query.Where("t.outlet_id = ?", outletId)
		}
	}

	if startDate != "" {
		query.Where("t.created_at > ?", startDate)

		if endDate != "" {
			query.Where("t.created_at < ?", endDate)
		}
	} else {
		if endDate != "" {
			query.Where("t.created_at < ?", endDate)
		}
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 5
	}

	query.Limit(int(limit))

	offset := (page - 1) * limit
	query.Offset(int(offset))

	query.Group("t.outlet_id")
	query.Count(&total)

	lastPage = math.Ceil(float64(total / limit))
	if err := query.Scan(&dataList).Error; err != nil {
		return dataList, total, page, lastPage, err
	}

	log.Info("result data list :", dataList)
	fmt.Println("datalist : ", dataList)
	return dataList, total, page, lastPage, err
}
