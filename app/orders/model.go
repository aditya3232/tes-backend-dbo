package orders

import "time"

type Orders struct {
	ID          int        `gorm:"primary_key;column:id" json:"id"`
	CustomerID  *int       `gorm:"column:customer_id" json:"customer_id"`
	TotalAmount int        `gorm:"column:total_amount" json:"total_amount"`
	Status      string     `gorm:"column:status" json:"status"`
	PaymentType string     `gorm:"column:payment_type" json:"payment_type"`
	CreatedAt   *time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
}

func (m *Orders) TableName() string {
	return "orders"
}
