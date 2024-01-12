package orders

type OrdersInput struct {
	CustomerID  *int   `form:"customer_id" binding:"required"`
	TotalAmount int    `form:"total_amount" binding:"required"`
	Status      string `form:"status" binding:"required"`
	PaymentType string `form:"payment_type" binding:"required"`
}

type OrdersUpdateInput struct {
	ID          int    `form:"id"` // buat update
	CustomerID  *int   `form:"customer_id"`
	TotalAmount int    `form:"total_amount"`
	Status      string `form:"status"`
	PaymentType string `form:"payment_type"`
}

type OrdersGetOneByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
