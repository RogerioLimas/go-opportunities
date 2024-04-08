package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// Listen and Serve in 0.0.0.0:8080
	router.Run(":8080")
}
