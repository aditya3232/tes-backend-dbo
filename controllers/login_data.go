package controllers

import (
	"net/http"

	login_data_app "github.com/aditya3232/tes-backend-dbo/app/login_data"
	users_app "github.com/aditya3232/tes-backend-dbo/app/users"
	"github.com/aditya3232/tes-backend-dbo/constant"
	"github.com/aditya3232/tes-backend-dbo/helper"
	log_function "github.com/aditya3232/tes-backend-dbo/log"
	"github.com/gin-gonic/gin"
)

type LoginDataController struct {
	loginDataService login_data_app.Service
}

func NewLoginDataController(loginDataService login_data_app.Service) *LoginDataController {
	return &LoginDataController{loginDataService}
}

func (h *LoginDataController) GetLoginData(c *gin.Context) {
	userID := c.MustGet("currentUser").(users_app.Users).ID

	loginData, err := h.loginDataService.GetLoginData(userID)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, users_app.UsersGetFormat(loginData))
	c.JSON(response.Meta.Code, response)

}
