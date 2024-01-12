package customers

type CustomersInput struct {
	UserID  *int   `form:"user_id"`
	Name    string `form:"name" binding:"required"`
	Email   string `form:"email" binding:"required"`
	Phone   string `form:"phone" binding:"required"`
	Street  string `form:"street" binding:"required"`
	ZipCode int    `form:"zip_code" binding:"required"`
	City    string `form:"city" binding:"required"`
	Country string `form:"country" binding:"required"`
}

type CustomersUpdateInput struct {
	ID      int    `form:"id"` // buat update
	UserID  *int   `form:"user_id"`
	Name    string `form:"name"`
	Email   string `form:"email"`
	Phone   string `form:"phone"`
	Street  string `form:"street"`
	ZipCode int    `form:"zip_code"`
	City    string `form:"city"`
	Country string `form:"country"`
}

type CustomersGetOneByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
