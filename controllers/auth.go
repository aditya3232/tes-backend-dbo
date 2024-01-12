package controllers

import (
	"net/http"
	"time"

	"github.com/aditya3232/tes-backend-dbo/constant"

	auth_app "github.com/aditya3232/tes-backend-dbo/app/auth"
	"github.com/aditya3232/tes-backend-dbo/helper"
	log_function "github.com/aditya3232/tes-backend-dbo/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct {
	authService auth_app.Service
}

func NewAuthController(authService auth_app.Service) *AuthController {
	return &AuthController{authService}
}

func (h *AuthController) Login(c *gin.Context) {
	var input auth_app.LoginInput

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.LoginFailed
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	token, err := h.authService.Login(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.LoginFailed
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.LoginSuccess
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	newToken := token.RememberToken
	expires := time.Now().AddDate(0, 0, 30)
	response := helper.APIResponse(message, http.StatusOK, auth_app.LoginFormat(newToken, expires))
	c.JSON(response.Meta.Code, response)
}

func (h *AuthController) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == " " {
		endpoint := c.Request.URL.Path
		message := constant.LoginFailed
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := constant.TokenNotValid
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	err := h.authService.Logout(token)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.LogoutFailed
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.LogoutSuccess
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, nil)
	c.JSON(response.Meta.Code, response)
}
