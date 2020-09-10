package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication : we start the application
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
