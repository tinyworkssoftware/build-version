package custom_errors

import (
	"build-version/model/response"
	"errors"
	"fmt"
	"time"
)

type DatabaseError struct {
	CorrelationId string
	ErrorType string
	Err error
}

type ErrorInterface interface {
	ToErrorResponse()
	Error() error
}

func (e DatabaseError) ToErrorResponse() *response.Error {
	return &response.Error{
		ErrorMessage:  fmt.Sprintf("Error [%s]: %s", e.ErrorType, e.Err.Error()),
		CorrelationId: e.CorrelationId,
		TransactionTs: time.Time{},
	}
}

func NewDatabaseError(correlationId string, errorType string, description string) *DatabaseError {
	return &DatabaseError{
		CorrelationId: correlationId,
		ErrorType:     errorType,
		Err:           errors.New(description),
	}
}