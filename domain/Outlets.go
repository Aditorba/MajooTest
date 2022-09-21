package domain

import "time"

type Outlets struct {
	Id         int64     `gorm:"type:bigint(20); primaryKey; not null; autoIncrement;" json:"id"`
	MerchantId int64     `gorm:"type:bigint(20); not null;" json:"merchantId"`
	OutletName string    `gorm:"type:varchar(40); not null;" json:"outletName"`
	CreatedAt  time.Time `gorm:"type:timestamp; not null; default:current_timestamp;" json:"createdAt"`
	CreatedBy  int64     `gorm:"type:bigint(20); not null;" json:"createdBy"`
	UpdateAt   time.Time `gorm:"type:timestamp; not null; default:current_timestamp;" json:"updateAt"`
	UpdatedBy  int64     `gorm:"type:bigint(20); not null;" json:"updatedBy"`
	table      string    `gorm:"-"`
}

func (p Outlets) TableName() string {
	if p.table != "" {
		return p.table
	}
	return "Outlets" // default table name
}
