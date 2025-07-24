package main

import (
	"fmt"

	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controller := controllers.NewController()

	controller.SetupRoutes(router)

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Health check available at: http://localhost:8080/health")

	router.Run(port)
}
