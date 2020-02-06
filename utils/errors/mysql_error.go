package errors

import (
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no row in result set"
)

//ParseError . . .
func ParseError(err error) *RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return NewNotFoundError("no record matching the given id")
		}
		return NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return NewBadRequest("invalid data")
	}
	return NewInternalServerError("error processing request")

}
