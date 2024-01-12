package users

type UsersInput struct {
	RoleID   *int   `form:"role_id" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UsersUpdateInput struct {
	ID            int    `form:"id"` // buat update
	RoleID        *int   `form:"role_id"`
	Password      string `form:"password"`
	RememberToken string `form:"remember_token"` // dari login
}

type UsersGetOneByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
