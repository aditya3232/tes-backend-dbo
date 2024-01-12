package users

import "time"

type Users struct {
	ID            int        `gorm:"primary_key;column:id" json:"id"`
	Username      string     `gorm:"unique;column:username" json:"username"`
	Password      string     `gorm:"column:password" json:"password"`
	RememberToken string     `gorm:"column:remember_token" json:"remember_token"`
	CreatedAt     *time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
}

func (m *Users) TableName() string {
	return "users"
}
