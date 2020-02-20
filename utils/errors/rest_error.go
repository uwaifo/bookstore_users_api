package errors

import (
	"errors"
	"net/http"
)

//RestErr . . .
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json: "error"`
}

//NewError . . .
func NewError(msg string) error {
	return errors.New(msg)

}

//NewBadRequest  . . .
func NewBadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}

}

//NewNotFoundError for queries that yield no result
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not found",
	}

}

//NewInternalServerError ...
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal server error",
	}
}
