package bcerrors

import (
	"fmt"

	"github.com/google/uuid"
)

//Error to wrap generic error with more meaningful one
type Error struct {
	id              string
	errType         string
	errCode         int
	rootCause       error
	externalMessage string
	message         string
}

//NewError constructs error
func NewError(message string, errType int) *Error {
	return &Error{
		id:      uuid.New().String(),
		message: message,
		errType: getTypeString(errType),
		errCode: getTypeCode(errType),
	}
}

//Error string representation of error
func (err *Error) Error() string {

	var rootCause string = ""
	if err.rootCause != nil {
		rootCause = err.rootCause.Error()
	}
	return fmt.Sprintf(
		"error_id: %s, type: %s, message: %s, root_cause: %s, external_message: %s",
		err.id,
		err.errType,
		err.message,
		rootCause,
		err.externalMessage,
	)
}

//WithExternalMessage adds external message field
func (err *Error) WithExternalMessage(msg string) *Error {
	err.externalMessage = msg
	return err
}

//WithRootCause adds root cause error field
func (err *Error) WithRootCause(cause error) *Error {
	err.rootCause = cause
	return err
}
