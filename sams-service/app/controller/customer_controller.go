package controller

import (
	"net/http"
	customer_service "sams/app/service"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Service *customer_service.CustomerService
}

func NewCustomerController(service *customer_service.CustomerService) *CustomerController {
	return &CustomerController{Service: service}
}

func (c *CustomerController) GetAllCustomers(ctx *gin.Context) {
	customers, err := c.Service.GetAllCustomers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	response := gin.H{
		"code": 200,
		"data": customers,
	}

	ctx.JSON(http.StatusOK, response)
}
