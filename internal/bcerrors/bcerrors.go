package bcerrors

import (
	"encoding/json"
	"fmt"
	"strings"

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

	errMap := make(map[string]string)

	errMap["id"] = err.id
	errMap["type"] = err.errType
	errMap["message"] = err.message

	if err.rootCause != nil {
		errMap["root_cause"] = err.rootCause.Error()
	}
	if err.externalMessage != "" {
		errMap["external_message"] = err.externalMessage
	}
	jsonString, _ := json.Marshal(errMap)

	fmt.Println(strings.Replace(string(jsonString), "\\", "", -1))
	return strings.Replace(string(jsonString), "\\", "", -1)
}
func (err *Error) WithExternalMessage(msg string) *Error {
	err.externalMessage = msg
	return err
}
func (err *Error) WithRootCause(cause error) *Error {
	err.rootCause = cause
	return err
}
