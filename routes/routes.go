package routes

import (
	"github.com/aditya3232/tes-backend-dbo/app/auth"
	"github.com/aditya3232/tes-backend-dbo/app/customers"
	"github.com/aditya3232/tes-backend-dbo/app/orders"
	"github.com/aditya3232/tes-backend-dbo/app/users"
	"github.com/aditya3232/tes-backend-dbo/connection"
	"github.com/aditya3232/tes-backend-dbo/controllers"
	"github.com/aditya3232/tes-backend-dbo/middleware"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	usersRepository := users.NewRepository(connection.DatabaseMysql())
	customersRepository := customers.NewRepository(connection.DatabaseMysql())
	ordersRepository := orders.NewRepository(connection.DatabaseMysql())

	// Initialize services
	usersService := users.NewService(usersRepository)
	customersService := customers.NewService(customersRepository, usersRepository)
	ordersService := orders.NewService(ordersRepository)
	authService := auth.NewService(usersRepository)

	// Initialize handlers
	UsersController := controllers.NewUsersController(usersService)
	CustomersController := controllers.NewCustomersController(customersService)
	OrdersController := controllers.NewOrdersController(ordersService)
	AuthController := controllers.NewAuthController(authService)

	// Configure routes
	api := router.Group("/api/tesbedbo/v1")

	usersRoutes := api.Group("/users")
	customersRoutes := api.Group("/customers", middleware.AuthMiddleware(usersService))
	ordersRoutes := api.Group("/orders", middleware.AuthMiddleware(usersService))
	authRoutes := api.Group("/auth")

	configureUsersRoutes(usersRoutes, UsersController)
	configureCustomersRoutes(customersRoutes, CustomersController)
	configureOrdersRoutes(ordersRoutes, OrdersController)
	configureAuthRoutes(authRoutes, AuthController)
}

func configureUsersRoutes(group *gin.RouterGroup, controller *controllers.UsersController) {
	group.GET("/getwithpaginateandsearch", controller.GetAll)
	group.POST("/insert", controller.Create)
	group.GET("/getdetail/:id", controller.GetOne)
	group.PUT("/update/:id", controller.Update)
	group.DELETE("/delete/:id", controller.Delete)
}

func configureCustomersRoutes(group *gin.RouterGroup, controller *controllers.CustomersController) {
	group.GET("/getwithpaginateandsearch", controller.GetAll)
	group.POST("/insert", controller.Create)
	group.GET("/getdetail/:id", controller.GetOne)
	group.PUT("/update/:id", controller.Update)
	group.DELETE("/delete/:id", controller.Delete)
}

func configureOrdersRoutes(group *gin.RouterGroup, controller *controllers.OrdersController) {
	group.GET("/getwithpaginateandsearch", controller.GetAll)
	group.POST("/insert", controller.Create)
	group.GET("/getdetail/:id", controller.GetOne)
	group.PUT("/update/:id", controller.Update)
	group.DELETE("/delete/:id", controller.Delete)
}

func configureAuthRoutes(group *gin.RouterGroup, controller *controllers.AuthController) {
	group.POST("/login", controller.Login)
	group.POST("/logout", controller.Logout)
}
