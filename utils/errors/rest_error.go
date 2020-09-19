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

// BadRequestError : return bad request rest error
func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "BAD_REQUEST",
	}
}

// NotFoundError : return not found rest error
func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "NOT_FOUND",
	}
}

// InternalServerError : return internal server rest error
func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "INTERNAL_SERVER_ERROR",
	}
}
