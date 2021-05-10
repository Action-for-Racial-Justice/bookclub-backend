package bcerrors_test

import (
	"testing"

	error "github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/stretchr/testify/assert"
)

const (
	expectedInternal   string = "type: Internal Server Error, message: This is not an internal error it's a test, root_cause: , external_message: "
	expectedValidation string = "type: Validation Error, message: This is not a validation error it's a test, root_cause: , external_message: "
	expectedDecode     string = "type: JSON Decode Error, message: This is not a decode error it's a test, root_cause: , external_message: "
)

func TestNewErrors(t *testing.T) {

	errInternal := error.NewError("This is not an internal error it's a test", error.InternalError)
	errValidation := error.NewError("This is not a validation error it's a test", error.ValidationError)
	errDecode := error.NewError("This is not a decode error it's a test", error.DecodeError)

	assert.NotNil(t, errInternal)
	assert.NotNil(t, errValidation)
	assert.NotNil(t, errDecode)

	assert.Equal(t, expectedInternal, errInternal.Error()[48:])
	assert.Equal(t, expectedValidation, errValidation.Error()[48:])
	assert.Equal(t, expectedDecode, errDecode.Error()[48:])
}
