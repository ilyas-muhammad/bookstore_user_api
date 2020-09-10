package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser : to create new user
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// GetUser : to get user by id
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
