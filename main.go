package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// creating router
	// r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())
	// handle cors
	r.Use(HandleCORS())

	// register routes
	r.GET("/api/classify-number", HandleClassifyNumber)

	// starting the server
	log.Println("server starts at :8080")
	if err := r.Run(); err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}
}
