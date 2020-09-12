package users

import (
	"strings"

	"petaniweb.com/rest/v1/bookstore_user_api/utils/errors"
)

// User : user type
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate : user method to validate persitence data
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("Invalid email address")
	}

	return nil
}
