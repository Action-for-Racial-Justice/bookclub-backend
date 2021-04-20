//go:generate mockgen -package=mocks -destination=../mocks/validator.go github.com/Action-for-Racial-Justice/bookclub-backend/internal/validator Validator
package validator

import (
	"log"
	"strconv"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/google/wire"
)

var Module = wire.NewSet(
	New,
)

type Validator interface {
	ValidateUserID(userID string) error
}

type BCValidator struct {
}

func New() *BCValidator {
	return &BCValidator{}
}

const (
	genericErrMessage string = "cannot validate user_id field"
)

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
