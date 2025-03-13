package main

import (
	"fmt"
	"log"
	"net/http"

	"farzbook.com/m/handler"
	"farzbook.com/m/router"

	"farzbook.com/m/config"
)

func main() {
	// Load server configuration
	cfg := config.Load()

	// Initialize handlers
	farzbookHandler := handler.NewFarzbookHandler()

	// Setup router
	r := router.New()
	r.RegisterRoutes(farzbookHandler)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Server starting on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, r.Handler()))
}
