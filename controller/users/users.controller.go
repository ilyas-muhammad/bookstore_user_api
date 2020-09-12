package users

import (
	"net/http"
	"strconv"

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
	userID, errID := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if errID != nil {
		err := errors.BadRequestError("User ID should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if errID != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
