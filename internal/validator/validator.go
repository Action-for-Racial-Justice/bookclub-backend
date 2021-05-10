//go:generate mockgen -package=mocks -destination=../mocks/validator.go github.com/Action-for-Racial-Justice/bookclub-backend/internal/validator Validator

package validator

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/google/wire"
)

const (
	genericErrMessage string = "cannot validate user_id field"
	tokenErrMessage   string = "token cannot be validated"
)

// Module to associate wire bindings
var Module = wire.NewSet(
	New,
)

//Validator interface
type Validator interface {
	ValidateUserID(string) error
	ValidateSSOToken(string) error
}

//BCValidator short for book club validator
type BCValidator struct {
}

//New .. constructor
func New() *BCValidator {
	return &BCValidator{}
}

//ValidateUserID validates a userID field
func (bcv *BCValidator) ValidateUserID(userID string) error {

	// checks to see if we cast userID to integer type
	if _, err := strconv.Atoi(userID); err != nil {
		log.Println(err.Error())
		return bcerrors.NewError(genericErrMessage, bcerrors.ValidationError).
			WithExternalMessage("user_id must be of integer type").
			WithRootCause(err)
	}
	return nil
}

//ValidateSSOToken validates sso token format
func (bcv *BCValidator) ValidateSSOToken(token string) error {
	if tokenLength := len(token); tokenLength != 36 {
		return bcerrors.NewError(genericErrMessage, bcerrors.ValidationError).
			WithExternalMessage(fmt.Sprintf("sso token must be of length 36, got: %d", tokenLength))
	}

	return nil
}
