package users

import "time"

type Users struct {
	ID            int        `gorm:"primary_key;column:id" json:"id"`
	RoleID        *int       `gorm:"column:role_id" json:"role_id"`
	Username      string     `gorm:"column:username" json:"username"`
	Password      string     `gorm:"column:password" json:"password"`
	RememberToken string     `gorm:"column:remember_token" json:"remember_token"`
	CreatedAt     *time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTTime" json:"updated_at"`
}

func (m *Users) TableName() string {
	return "users"
}
