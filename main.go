package main

import (
	"fmt"
	"log"

	"farzbook.com/m/config"
	"farzbook.com/m/handler"
	"github.com/gin-gonic/gin"
)

func main() {
        cfg := config.Load()

        farzbookHandler := handler.NewFarzbookHandler()

        r := gin.Default()

        r.GET("/", farzbookHandler.Home)
        r.GET("/login", farzbookHandler.Login)
        r.GET("/profile", farzbookHandler.Profile)
        r.GET("/logout", farzbookHandler.Logout)
        r.GET("/add", farzbookHandler.Add)
        r.GET("/sub", farzbookHandler.Sub)

        addr := fmt.Sprintf(":%d", cfg.Port)
        log.Printf("Server starting on http://localhost%s", addr)
        log.Fatal(r.Run(addr))
}
