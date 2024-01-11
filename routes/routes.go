package routes

import (
	"github.com/aditya3232/tes-backend-dbo/app/users"
	"github.com/aditya3232/tes-backend-dbo/connection"
	"github.com/aditya3232/tes-backend-dbo/controllers"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	usersRepository := users.NewRepository(connection.DatabaseMysql())

	// Initialize services
	usersService := users.NewService(usersRepository)

	// Initialize handlers
	UsersController := controllers.NewUsersController(usersService)

	// Configure routes
	api := router.Group("/api/tesbebdo/v1")

	usersRoutes := api.Group("/users")

	configureUsersRoutes(usersRoutes, UsersController)
}

func configureUsersRoutes(group *gin.RouterGroup, controller *controllers.UsersController) {
	group.GET("/getwithpaginateandsearch", controller.GetAll)
	group.POST("/insert", controller.Create)
	group.GET("/getdetail/:id", controller.GetOne)
	group.PUT("/update/:id", controller.Update)
	group.DELETE("/delete/:id", controller.Delete)
}
