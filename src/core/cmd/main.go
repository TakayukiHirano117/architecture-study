package main

import (
	"fmt"

	// "github.com/TakayukiHirano117/architecture-study/config"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	// dbConfig := config.NewDBConfig()
	// fmt.Println(dbConfig)

	router := gin.Default()

	controller := controllers.NewController()

	controller.SetupRoutes(router)
	controllers.NewUserController(router)


	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Health check available at: http://localhost:8080/health")

	router.Run(port)
}
