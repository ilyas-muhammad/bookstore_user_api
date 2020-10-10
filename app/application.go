package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyas-muhammad/bookstore_user_api/logger"
)

var (
	router = gin.Default()
)

// StartApplication : we start the application
func StartApplication() {
	mapUrls()

	port := ":8080"
	logger.Info("Starting server...")
	err := router.Run(port)
	if err != nil {
		panic(err)
	}
}
