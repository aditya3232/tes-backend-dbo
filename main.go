package main

import (
	"fmt"

	"github.com/aditya3232/tes-backend-dbo/config"
	"github.com/aditya3232/tes-backend-dbo/helper"
	"github.com/aditya3232/tes-backend-dbo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	defer helper.RecoverPanic()

	router := gin.Default()
	if config.CONFIG.DEBUG == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	routes.Initialize(router)
	router.Run(fmt.Sprintf("%s:%s", config.CONFIG.APP_HOST, config.CONFIG.APP_PORT))
}
