package services

import (
	"petaniweb.com/rest/v1/bookstore_user_api/domain/users"
	"petaniweb.com/rest/v1/bookstore_user_api/utils/errors"
)

// GetUser : get user by id
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateUser : user creation business logic
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
