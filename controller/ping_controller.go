package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping request handler
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
