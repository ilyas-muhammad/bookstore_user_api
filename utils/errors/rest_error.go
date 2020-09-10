package errors

import (
	"net/http"
)

// RestErr : error type
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// BadRequestError : return rest error related to bad request
func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "BAD_REQUEST",
	}
}
