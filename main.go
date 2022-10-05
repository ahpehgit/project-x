package main

import (
	"log"
	"time"

	constants "github.com/ahpehgit/go-project-x/constants"
	controllers "github.com/ahpehgit/go-project-x/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()

	// Set run mode
	mode := gin.DebugMode

	env := viper.GetString("ENVIRONMENT")
	if env == "release" {
		log.Print("*** Running in Release mode ***")
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	// Set CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     constants.GetOrigins(),
		AllowMethods:     constants.GetAllowMethods(),
		AllowHeaders:     constants.GetAllowHeaders(),
		ExposeHeaders:    constants.GetExposedHeaders(),
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// Define endpoints
	mainController := controllers.NewMainController()

	router.GET(constants.ROOT, mainController.Root)
	router.SetTrustedProxies([]string{"localhost"})

	// Get host and port from env variables, if not default to config values
	host := viper.GetString("HOST")
	port := viper.GetString("PORT")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3003"
	}

	// Start server
	log.Printf("Server running in %s:%s", host, port)
	router.Run(host + ":" + port)
}
