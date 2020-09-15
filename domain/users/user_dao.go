package users

import (
	"fmt"

	"petaniweb.com/rest/v1/bookstore_user_api/utils/dateutils"

	"petaniweb.com/rest/v1/bookstore_user_api/utils/errors"
)

// mocking DB
var usersDB = make(map[int64]*User)

// Get : Method to get user by id
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]

	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("User %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Save : Method to create new user
func (user *User) Save() *errors.RestErr {
	if usersDB[user.ID] != nil {
		if usersDB[user.ID].Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("User %s already taken", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("User %d already exist", user.ID))
	}

	user.DateCreated = dateutils.GetNowString()

	usersDB[user.ID] = user
	return nil
}
