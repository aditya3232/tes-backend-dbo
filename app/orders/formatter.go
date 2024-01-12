package orders

import "time"

type OrdersGetFormatter struct {
	ID          int        `json:"id"`
	CustomerID  *int       `json:"customer_id"`
	TotalAmount int        `json:"total_amount"`
	Status      string     `json:"status"`
	PaymentType string     `json:"payment_type"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type OrdersCreateFormatter struct {
	ID          int        `json:"id"`
	CustomerID  *int       `json:"customer_id"`
	TotalAmount int        `json:"total_amount"`
	Status      string     `json:"status"`
	PaymentType string     `json:"payment_type"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type OrdersUpdateFormatter struct {
	ID          int        `json:"id"`
	CustomerID  *int       `json:"customer_id"`
	TotalAmount int        `json:"total_amount"`
	Status      string     `json:"status"`
	PaymentType string     `json:"payment_type"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func OrdersCreateFormat(orders Orders) OrdersCreateFormatter {
	var formatter OrdersCreateFormatter

	formatter.ID = orders.ID
	formatter.CustomerID = orders.CustomerID
	formatter.TotalAmount = orders.TotalAmount
	formatter.Status = orders.Status
	formatter.PaymentType = orders.PaymentType
	formatter.CreatedAt = orders.CreatedAt
	formatter.UpdatedAt = orders.UpdatedAt

	return formatter
}

func OrdersUpdateFormat(orders Orders) OrdersUpdateFormatter {
	var formatter OrdersUpdateFormatter

	formatter.ID = orders.ID
	formatter.CustomerID = orders.CustomerID
	formatter.TotalAmount = orders.TotalAmount
	formatter.Status = orders.Status
	formatter.PaymentType = orders.PaymentType
	formatter.CreatedAt = orders.CreatedAt
	formatter.UpdatedAt = orders.UpdatedAt

	return formatter
}

func OrdersGetFormat(orders Orders) OrdersGetFormatter {
	var formatter OrdersGetFormatter

	formatter.ID = orders.ID
	formatter.CustomerID = orders.CustomerID
	formatter.TotalAmount = orders.TotalAmount
	formatter.Status = orders.Status
	formatter.PaymentType = orders.PaymentType
	formatter.CreatedAt = orders.CreatedAt
	formatter.UpdatedAt = orders.UpdatedAt

	return formatter
}

func OrdersGetAllFormat(orders []Orders) []OrdersGetFormatter {
	formatter := []OrdersGetFormatter{}

	for _, order := range orders {
		OrdersGetFormatter := OrdersGetFormat(order)
		formatter = append(formatter, OrdersGetFormatter)
	}

	return formatter
}
