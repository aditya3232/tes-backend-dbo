package controllers

import (
	"net/http"

	users_app "github.com/aditya3232/tes-backend-dbo/app/users"
	"github.com/aditya3232/tes-backend-dbo/constant"
	"github.com/aditya3232/tes-backend-dbo/helper"
	log_function "github.com/aditya3232/tes-backend-dbo/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UsersController struct {
	usersService users_app.Service
}

func NewUsersController(usersService users_app.Service) *UsersController {
	return &UsersController{usersService}
}

func (h *UsersController) GetAll(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, users_app.Users{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	users, page, err := h.usersService.GetAll(filter, page, sort)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(message, http.StatusNotFound, helper.Pagination{}, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(users) == 0 {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := "Users not found"
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(message, http.StatusNotFound, helper.Pagination{}, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIDataTableResponse(message, http.StatusOK, page, users_app.UsersGetAllFormat(users))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersController) GetOne(c *gin.Context) {
	var input users_app.UsersGetOneByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	user, err := h.usersService.GetOne(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusNotFound, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, users_app.UsersGetFormat(user))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersController) Create(c *gin.Context) {
	var input users_app.UsersInput

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	newUser, err := h.usersService.Create(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessCreateData
	infoCode := http.StatusCreated
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusCreated, users_app.UsersCreateFormat(newUser))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersController) Update(c *gin.Context) {
	var id users_app.UsersGetOneByIdInput
	var input users_app.UsersUpdateInput

	err := c.ShouldBindUri(&id)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	input.ID = id.ID

	err = c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	newUser, err := h.usersService.Update(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.FailedUpdateData
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessUpdateData
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, users_app.UsersUpdateFormat(newUser))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersController) Delete(c *gin.Context) {
	var input users_app.UsersGetOneByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	err = h.usersService.Delete(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.FailedDeleteData
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessDeleteData
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, nil)
	c.JSON(response.Meta.Code, response)
}
