package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/config"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/middlewares"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.Use(middlewares.ErrorHandlingMiddleware())

	conn, err := rdb.NewConnection(config.NewDBConfig())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	router.Use(middlewares.DBMiddleware(conn))

	controller := controllers.NewController()
	controller.SetupRoutes(router)

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)

	if err := router.Run(port); err != nil {
		panic(err)
	}
}
