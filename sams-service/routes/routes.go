package routes

import (
	"sams/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.Use(cors.Default())
	r.Use(middleware.CorsMiddleware())
	RegisterCustomerRoutes(r, db)
}
