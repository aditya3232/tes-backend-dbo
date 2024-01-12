package middleware

import (
	"net/http"

	users_app "github.com/aditya3232/tes-backend-dbo/app/users"
	"github.com/aditya3232/tes-backend-dbo/helper"
	"github.com/aditya3232/tes-backend-dbo/library/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(usersService users_app.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		userID, err := jwt.GetUserIDFromToken(authHeader)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// var
		var UsersGetOneByIdInput users_app.UsersGetOneByIdInput
		UsersGetOneByIdInput.ID = userID

		user, err := usersService.GetOne(UsersGetOneByIdInput)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if user.RememberToken != authHeader {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

		c.Next()
	}
}
