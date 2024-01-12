package auth

type LoginInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LogoutInput struct {
	Token string `header:"Authorization"`
}
