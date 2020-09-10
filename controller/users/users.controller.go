package users

import (
	"net/http"

	"petaniweb.com/rest/v1/bookstore_user_api/utils/errors"

	"petaniweb.com/rest/v1/bookstore_user_api/services"

	"petaniweb.com/rest/v1/bookstore_user_api/domain/users"

	"github.com/gin-gonic/gin"
)

// CreateUser : to create new user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid json body request")

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser : to get user by id
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
