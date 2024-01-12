package customers

import (
	"time"
)

type CustomersGetFormatter struct {
	ID        int        `json:"id"`
	UserID    *int       `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Street    string     `json:"street"`
	ZipCode   int        `json:"zip_code"`
	City      string     `json:"city"`
	Country   string     `json:"country"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CustomersCreateFormatter struct {
	ID        int        `json:"id"`
	UserID    *int       `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Street    string     `json:"street"`
	ZipCode   int        `json:"zip_code"`
	City      string     `json:"city"`
	Country   string     `json:"country"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CustomersUpdateFormatter struct {
	ID        int        `json:"id"`
	UserID    *int       `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Street    string     `json:"street"`
	ZipCode   int        `json:"zip_code"`
	City      string     `json:"city"`
	Country   string     `json:"country"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func CustomersCreateFormat(customers Customers) CustomersCreateFormatter {
	var formatter CustomersCreateFormatter

	formatter.ID = customers.ID
	formatter.UserID = customers.UserID
	formatter.Name = customers.Name
	formatter.Email = customers.Email
	formatter.Phone = customers.Phone
	formatter.Street = customers.Street
	formatter.ZipCode = customers.ZipCode
	formatter.City = customers.City
	formatter.Country = customers.Country
	formatter.CreatedAt = customers.CreatedAt
	formatter.UpdatedAt = customers.UpdatedAt

	return formatter
}

func CustomersUpdateFormat(customers Customers) CustomersUpdateFormatter {
	var formatter CustomersUpdateFormatter

	formatter.ID = customers.ID
	formatter.UserID = customers.UserID
	formatter.Name = customers.Name
	formatter.Email = customers.Email
	formatter.Phone = customers.Phone
	formatter.Street = customers.Street
	formatter.ZipCode = customers.ZipCode
	formatter.City = customers.City
	formatter.Country = customers.Country
	formatter.CreatedAt = customers.CreatedAt
	formatter.UpdatedAt = customers.UpdatedAt

	return formatter
}

func CustomersGetFormat(customers Customers) CustomersGetFormatter {
	var formatter CustomersGetFormatter

	formatter.ID = customers.ID
	formatter.UserID = customers.UserID
	formatter.Name = customers.Name
	formatter.Email = customers.Email
	formatter.Phone = customers.Phone
	formatter.Street = customers.Street
	formatter.ZipCode = customers.ZipCode
	formatter.City = customers.City
	formatter.Country = customers.Country
	formatter.CreatedAt = customers.CreatedAt
	formatter.UpdatedAt = customers.UpdatedAt

	return formatter
}

func CustomersGetAllFormat(customers []Customers) []CustomersGetFormatter {
	formatter := []CustomersGetFormatter{}

	for _, customer := range customers {
		CustomersGetFormatter := CustomersGetFormat(customer)
		formatter = append(formatter, CustomersGetFormatter)
	}

	return formatter
}
