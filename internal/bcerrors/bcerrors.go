package bcerrors

import (
	"fmt"

	"github.com/google/uuid"
)

type Error struct {
	id              string
	errType         string
	errCode         int
	rootCause       error
	externalMessage string
	message         string
}

func NewError(message string, errType int) *Error {
	return &Error{
		id:      uuid.New().String(),
		message: message,
		errType: getTypeString(errType),
		errCode: getTypeCode(errType),
	}
}

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

func (err *Error) WithExternalMessage(msg string) *Error {
	err.externalMessage = msg
	return err
}
func (err *Error) WithRootCause(cause error) *Error {
	err.rootCause = cause
	return err
}
