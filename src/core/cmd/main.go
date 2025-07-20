package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers"
)

func main() {
	controller := controllers.NewController()

	mux := http.NewServeMux()

	controller.SetupRoutes(mux)

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Health check available at: http://localhost:8080/health")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
