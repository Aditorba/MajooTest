package domain

import "time"

type Users struct {
	Id        int64     `gorm:"type:bigint(20); primaryKey; not null; autoIncrement;" json:"id"`
	Name      string    `gorm:"type:varchar(45); default:null;" json:"name"`
	UserName  string    `gorm:"type:varchar(45); default:null;" json:"userName"`
	Password  string    `gorm:"type:varchar(225);" json:"password"`
	Token     string    `gorm:"type:varchar(225);" json:"token"`
	CreatedAt time.Time `gorm:"type:timestamp; not null; default:current_timestamp;" json:"createdAt"`
	CreatedBy int64     `gorm:"type:bigint(20); not null;" json:"createdBy"`
	UpdateAt  time.Time `gorm:"type:timestamp; not null; default:current_timestamp;" json:"updateAt"`
	UpdatedBy int64     `gorm:"type:bigint(20); not null;" json:"updatedBy"`
	table     string    `gorm:"-"`
}

func (p Users) TableName() string {
	if p.table != "" {
		return p.table
	}
	return "Users" // default table name
}
