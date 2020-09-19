package users

import (
	"fmt"
	"github.com/ilyas-muhammad/bookstore_user_api/datasources/mysql/users_db"
	"github.com/ilyas-muhammad/bookstore_user_api/utils/dateutils"
	"strings"

	"github.com/ilyas-muhammad/bookstore_user_api/utils/errors"
)

const (
	indexUniqueEmail = "unique_email"
	noRowsInResult   = "no rows in result set"
	insertQuery      = "INSERT INTO users(first_name, last_name, email) VALUES (?,?,?);"
	getByIdQuery     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

// Get : Method to get user by id
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(getByIdQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), noRowsInResult) {
			return errors.NotFoundError(fmt.Sprintf("User %d not found", user.ID))
		}
		return errors.InternalServerError(
			fmt.Sprintf("Error when trying to get user: %s", err.Error()),
		)
	}

	return nil
}

// Save : Method to create new user
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(insertQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.BadRequestError(
				fmt.Sprintf("email: %s already taken", user.Email),
			)
		}
		return errors.InternalServerError(
			fmt.Sprintf("Error when trying save users, %s", err.Error()),
		)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError(
			fmt.Sprintf("Error when trying save users, %s", err.Error()),
		)
	}

	user.ID = userID
	user.DateCreated = dateutils.GetNowString()
	return nil
}
