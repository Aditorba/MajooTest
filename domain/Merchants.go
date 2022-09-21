package domain

import "time"

type Merchants struct {
	Id           int64     `gorm:"type:bigint(20); primaryKey; not null; autoIncrement;" json:"id"`
	UserId       int32     `gorm:"type:int(40); not null;" json:"userId"`
	MerchantName string    `gorm:"type:varchar(40); not null;" json:"merchantName"`
	CreatedAt    time.Time `gorm:"type:timestamp; not null; default:current_timestamp;" json:"createdAt"`
	CreatedBy    int64     `gorm:"type:bigint(20); not null;" json:"createdBy"`
	UpdateAt     time.Time `gorm:"type:timestamp; not null; default:current_timestamp;" json:"updateAt"`
	UpdatedBy    int64     `gorm:"type:bigint(20); not null;" json:"updatedBy"`
	table        string    `gorm:"-"`
}

func (p Merchants) TableName() string {
	if p.table != "" {
		return p.table
	}
	return "Merchants" // default table name
}
