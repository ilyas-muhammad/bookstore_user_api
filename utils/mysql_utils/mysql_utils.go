package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/ilyas-muhammad/bookstore_user_api/utils/errors"
	"strings"
)

const (
	noRowsInResult = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsInResult) {
			return errors.NotFoundError("No record match given id")
		}
		return errors.InternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError("duplicated key")
	}

	return errors.InternalServerError("database error processing request")
}
