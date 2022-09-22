package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	router.SetTrustedProxies([]string{"localhost"})

	host := "localhost"
	port := "3001"

	// Start server
	log.Printf("Server running in %s:%s", host, port)
	router.Run(host + ":" + port)
}
