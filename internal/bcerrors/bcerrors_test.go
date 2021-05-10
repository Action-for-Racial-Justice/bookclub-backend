package bcerrors_test

import (
	"errors"
	"testing"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/stretchr/testify/assert"
)

const (
	expectedInternal       string = "type: Internal Server Error, message: This is not an internal error it's a test, root_cause: , external_message: "
	expectedInternalExRoot string = "type: Internal Server Error, message: This is not an internal error it's a test, root_cause: internal root error, external_message: internal external message"

	expectedValidation       string = "type: Validation Error, message: This is not a validation error it's a test, root_cause: , external_message: "
	expectedValidationExRoot string = "type: Validation Error, message: This is not a validation error it's a test, root_cause: validation root error, external_message: validation external message"

	expectedDecode       string = "type: JSON Decode Error, message: This is not a decode error it's a test, root_cause: , external_message: "
	expectedDecodeExRoot string = "type: JSON Decode Error, message: This is not a decode error it's a test, root_cause: decode root error, external_message: decode external message"
)

func TestNewErrors(t *testing.T) {

	errInternal := bcerrors.NewError("This is not an internal error it's a test", bcerrors.InternalError)
	errValidation := bcerrors.NewError("This is not a validation error it's a test", bcerrors.ValidationError)
	errDecode := bcerrors.NewError("This is not a decode error it's a test", bcerrors.DecodeError)

	assert.NotNil(t, errInternal)
	assert.NotNil(t, errValidation)
	assert.NotNil(t, errDecode)

	assert.Equal(t, expectedInternal, errInternal.Error()[48:])
	assert.Equal(t, expectedValidation, errValidation.Error()[48:])
	assert.Equal(t, expectedDecode, errDecode.Error()[48:])

	errInternal.WithExternalMessage("internal external message")
	assert.Equal(t, expectedInternalExRoot, errInternal.WithRootCause(errors.New("internal root error")).Error()[48:])

	errValidation.WithExternalMessage("validation external message")
	assert.Equal(t, expectedValidationExRoot, errValidation.WithRootCause(errors.New("validation root error")).Error()[48:])

	errDecode.WithExternalMessage("decode external message")
	assert.Equal(t, expectedDecodeExRoot, errDecode.WithRootCause(errors.New("decode root error")).Error()[48:])
}
