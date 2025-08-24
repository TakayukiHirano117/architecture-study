package main

import (
	"fmt"

	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.Use(middlewares.ErrorHandlingMiddleware())
	router.Use(middlewares.DBMiddleware())

	controller := controllers.NewController()
	controller.SetupRoutes(router)

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)

	router.Run(port)
}
