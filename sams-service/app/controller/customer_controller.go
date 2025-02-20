package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sams/app/service"
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

	ctx.JSON(http.StatusOK, customers)
}
