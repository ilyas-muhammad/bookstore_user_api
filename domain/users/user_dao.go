package users

import (
	"fmt"
	"github.com/ilyas-muhammad/bookstore_user_api/datasources/mysql/users_db"
	"github.com/ilyas-muhammad/bookstore_user_api/logger"
	"github.com/ilyas-muhammad/bookstore_user_api/utils/dateutils"
	"github.com/ilyas-muhammad/bookstore_user_api/utils/errors"
	"github.com/ilyas-muhammad/bookstore_user_api/utils/mysql_utils"
)

const (
	insertQuery       = "INSERT INTO users(first_name, last_name, email, status, password) VALUES (?,?,?,?,?);"
	getByIdQuery      = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	updateByIdQuery   = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	deleteByIdQuery   = "DELETE FROM users WHERE id=?;"
	findByStatusQuery = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

// Get : Method to get user by id
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(getByIdQuery)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.InternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

// Save : Method to create new user
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(insertQuery)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.InternalServerError("Database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user", err)
		return errors.InternalServerError("Database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying get last insert id after creating a user", err)
		return errors.InternalServerError("Database error")
	}

	user.ID = userID
	user.DateCreated = dateutils.GetNowString()
	return nil
}

// Update : Method to update existing user
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(updateByIdQuery)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.InternalServerError("Database error")
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		logger.Error("error when trying to update user", err)
		return errors.InternalServerError("Database error")
	}

	return nil
}

// Delete : Method to delete user
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(deleteByIdQuery)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.InternalServerError("Database error")
	}
	defer stmt.Close()

	if _, deleteErr := stmt.Exec(user.ID); deleteErr != nil {
		logger.Error("error when trying to delete user", err)
		return errors.InternalServerError("Database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(findByStatusQuery)
	if err != nil {
		logger.Error("error when trying to prepare find by status user statement", err)
		return nil, errors.InternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find by status user", err)
		return nil, errors.InternalServerError("Database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.InternalServerError("Database error")
		}
		results = append(results, user)
	}
	if len(results) < 1 {
		return nil, errors.NotFoundError(fmt.Sprintf("no users matching status: %s", status))
	}
	return results, nil
}
