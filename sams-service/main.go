package main

import (
	"github.com/gin-gonic/gin"
	"sams/config"
	"sams/routes"
)

func main() {
	router := gin.Default()

	db := config.InitPgDatabase()
	//r.Use(middleware.Logger())
	//r.Use(middleware.Recovery())

	routes.RegisterRoutes(router, db)

	if err := router.Run(":8080"); err != nil {
		panic(err)
		//log.Fatalf( "failed to run server: %v", err)
	}
}
