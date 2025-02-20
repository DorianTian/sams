package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sams/app/controller"
	"sams/app/service"
)

func RegisterCustomerRoutes(router *gin.Engine, db *gorm.DB) {
	customerService := customer_service.NewCustomerService(db)
	customerController := controller.CustomerController{Service: customerService}
	customers := router.Group("/customers")
	{
		customers.GET("", customerController.GetAllCustomers)
	}
}
