package customers

import (
	"time"
)

type Customers struct {
	ID        int        `gorm:"primary_key;column:id" json:"id"`
	UserID    *int       `gorm:"column:user_id" json:"user_id"`
	Name      string     `gorm:"column:name" json:"name"`
	Email     string     `gorm:"unique;column:email" json:"email"`
	Phone     string     `gorm:"column:phone" json:"phone"`
	Street    string     `gorm:"column:street" json:"street"`
	ZipCode   int        `gorm:"column:zip_code" json:"zip_code"`
	City      string     `gorm:"column:city" json:"city"`
	Country   string     `gorm:"column:country" json:"country"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
}

func (m *Customers) TableName() string {
	return "customers"
}
