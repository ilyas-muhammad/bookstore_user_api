package services

import (
	"petaniweb.com/rest/v1/bookstore_user_api/domain/users"
	"petaniweb.com/rest/v1/bookstore_user_api/utils/errors"
)

// CreateUser : user creation business logic
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
