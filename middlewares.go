package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusNoContent, gin.H{
				"error": http.StatusText(http.StatusNoContent),
			})
			// c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
