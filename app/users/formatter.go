package users

import "time"

type UsersGetFormatter struct {
	ID        int        `json:"id"`
	RoleID    *int       `json:"role_id"`
	Username  string     `json:"username"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UsersCreateFormatter struct {
	ID        int        `json:"id"`
	RoleID    *int       `json:"role_id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UsersUpdateFormatter struct {
	ID        int        `json:"id"`
	RoleID    *int       `json:"role_id"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func UsersCreateFormat(users Users) UsersCreateFormatter {
	var formatter UsersCreateFormatter

	formatter.ID = users.ID
	formatter.RoleID = users.RoleID
	formatter.Username = users.Username
	formatter.Password = users.Password
	formatter.CreatedAt = users.CreatedAt
	formatter.UpdatedAt = users.UpdatedAt

	return formatter
}

func UsersUpdateFormat(users Users) UsersUpdateFormatter {
	var formatter UsersUpdateFormatter

	formatter.ID = users.ID
	formatter.RoleID = users.RoleID
	formatter.Password = users.Password
	formatter.CreatedAt = users.CreatedAt
	formatter.UpdatedAt = users.UpdatedAt

	return formatter
}

func UsersGetFormat(users Users) UsersGetFormatter {
	var formatter UsersGetFormatter

	formatter.ID = users.ID
	formatter.RoleID = users.RoleID
	formatter.Username = users.Username
	formatter.CreatedAt = users.CreatedAt
	formatter.UpdatedAt = users.UpdatedAt

	return formatter
}

func UsersGetAllFormat(users []Users) []UsersGetFormatter {
	formatter := []UsersGetFormatter{}

	for _, user := range users {
		userGetFormatter := UsersGetFormat(user)
		formatter = append(formatter, userGetFormatter)
	}

	return formatter
}
