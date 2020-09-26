package users

import (
	"github.com/ilyas-muhammad/bookstore_user_api/datasources/mysql/users_db"
	"github.com/ilyas-muhammad/bookstore_user_api/utils/dateutils"
	"github.com/ilyas-muhammad/bookstore_user_api/utils/errors"
	"github.com/ilyas-muhammad/bookstore_user_api/utils/mysql_utils"
)

const (
	insertQuery     = "INSERT INTO users(first_name, last_name, email) VALUES (?,?,?);"
	getByIdQuery    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	updateByIdQuery = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

// Get : Method to get user by id
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(getByIdQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.ID = userID
	user.DateCreated = dateutils.GetNowString()
	return nil
}

// Update : Method to update existing user
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(updateByIdQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}
