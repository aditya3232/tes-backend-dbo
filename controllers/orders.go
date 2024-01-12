package controllers

import (
	"net/http"

	orders_app "github.com/aditya3232/tes-backend-dbo/app/orders"
	"github.com/aditya3232/tes-backend-dbo/constant"
	"github.com/aditya3232/tes-backend-dbo/helper"
	log_function "github.com/aditya3232/tes-backend-dbo/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrdersController struct {
	ordersService orders_app.Service
}

func NewOrdersController(ordersService orders_app.Service) *OrdersController {
	return &OrdersController{ordersService}
}

func (h *OrdersController) GetAll(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, orders_app.Orders{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	orders, page, err := h.ordersService.GetAll(filter, page, sort)
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

	if len(orders) == 0 {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := "Orders not found"
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

	response := helper.APIDataTableResponse(message, http.StatusOK, page, orders_app.OrdersGetAllFormat(orders))
	c.JSON(response.Meta.Code, response)

}

func (h *OrdersController) GetOne(c *gin.Context) {
	var input orders_app.OrdersGetOneByIdInput

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

	order, err := h.ordersService.GetOne(input)
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

	response := helper.APIResponse(message, http.StatusOK, orders_app.OrdersGetFormat(order))
	c.JSON(response.Meta.Code, response)
}

func (h *OrdersController) Create(c *gin.Context) {
	var input orders_app.OrdersInput

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

	newOrder, err := h.ordersService.Create(input)
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

	response := helper.APIResponse(message, http.StatusCreated, orders_app.OrdersCreateFormat(newOrder))
	c.JSON(response.Meta.Code, response)
}

func (h *OrdersController) Update(c *gin.Context) {
	var id orders_app.OrdersGetOneByIdInput
	var input orders_app.OrdersUpdateInput

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

	newOrder, err := h.ordersService.Update(input)
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

	response := helper.APIResponse(message, http.StatusOK, orders_app.OrdersUpdateFormat(newOrder))
	c.JSON(response.Meta.Code, response)
}

func (h *OrdersController) Delete(c *gin.Context) {
	var input orders_app.OrdersGetOneByIdInput

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

	err = h.ordersService.Delete(input)
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
