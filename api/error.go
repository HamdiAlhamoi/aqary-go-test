package api

import (
	"errors"

	"github.com/jackc/pgconn"
)

var (
	ErrNameAlreadyTaken = errors.New("name already taken")
	ErrPhoneNumberAlreadyTaken    = errors.New("phone number already taken")
	ErrAccessForbidden      = errors.New("access forbidden")
	ErrUserNotFound         = errors.New("user not found")
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewValidationError(err error) *Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = err.Error()
	return &e
}

func NewError(err error) *Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["message"] = err.Error()
	return &e
}

func convertToApiErr(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.ConstraintName {
		case "users_name_key":
			return ErrNameAlreadyTaken
		case "users_phone_number_key":
			return ErrPhoneNumberAlreadyTaken
		}
	}
	return nil
}

