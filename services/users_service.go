package services

import (
	"github.com/ilyas-muhammad/bookstore_user_api/domain/users"
	"github.com/ilyas-muhammad/bookstore_user_api/utils/errors"
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

// UpdateUser : user updating service
func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}
