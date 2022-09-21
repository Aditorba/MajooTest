package domain

import "time"

type Transactions struct {
	Id         int64     `gorm:"type:bigint(20); primaryKey; not null; autoIncrement;" json:"id"`
	MerchantId int64     `gorm:"type:bigint(20); not null;" json:"merchantId"`
	OutletId   int64     `gorm:"type:bigint(20); not null;" json:"outletId"`
	BillTotal  float64   `gorm:"type:double; not null;" json:"billTotal"`
	CreatedAt  time.Time `gorm:"type:timestamp; not null; default:current_timestamp;" json:"createdAt"`
	CreatedBy  int64     `gorm:"type:bigint(20); not null;" json:"createdBy"`
	UpdateAt   time.Time `gorm:"type:timestamp; not null; default:current_timestamp;" json:"updateAt"`
	UpdatedBy  int64     `gorm:"type:bigint(20); not null;" json:"updatedBy"`
	table      string    `gorm:"-"`
}

func (p Transactions) TableName() string {
	if p.table != "" {
		return p.table
	}
	return "Transactions" // default table name
}
