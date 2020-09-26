package users

import (
	"net/http"
	"strconv"

	"github.com/ilyas-muhammad/bookstore_user_api/utils/errors"

	"github.com/ilyas-muhammad/bookstore_user_api/services"

	"github.com/ilyas-muhammad/bookstore_user_api/domain/users"

	"github.com/gin-gonic/gin"
)

// ValidateUserIDParameter : DRY to get user by given request parameter id
func ValidateUserIDParameter(userIdParam string) (int64, *errors.RestErr) {
	userID, errID := strconv.ParseInt(userIdParam, 10, 64)
	if errID != nil {
		return 0, errors.BadRequestError("User ID should be a number")
	}

	return userID, nil
}

// Create : to create new user
func Create(c *gin.Context) {
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

// Get : to get user by id
func Get(c *gin.Context) {
	userID, errID := ValidateUserIDParameter(c.Param("user_id"))
	if errID != nil {
		c.JSON(errID.Status, errID)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update : to update user by given id
func Update(c *gin.Context) {
	userID, errID := ValidateUserIDParameter(c.Param("user_id"))
	if errID != nil {
		c.JSON(errID.Status, errID)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid json body request")

		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Delete : function to delete user by given id
func Delete(c *gin.Context) {
	userID, errID := ValidateUserIDParameter(c.Param("user_id"))
	if errID != nil {
		c.JSON(errID.Status, errID)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search : function to search user by given status
func Search(c *gin.Context) {
	status := c.Query("status")

	foundUsers, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, foundUsers)
}
