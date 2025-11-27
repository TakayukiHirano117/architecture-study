package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/config"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/middlewares"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

func main() {
	router := gin.Default()

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

	router.Run(port)

}
