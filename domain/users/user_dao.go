package users

import (
	"fmt"
	"github.com/ilyas-muhammad/bookstore_user_api/datasources/mysql/users_db"
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
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password)
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

// Delete : Method to delete user
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(deleteByIdQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, deleteErr := stmt.Exec(user.ID); deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(findByStatusQuery)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) < 1 {
		return nil, errors.NotFoundError(fmt.Sprintf("no users matching status: %s", status))
	}
	return results, nil
}
