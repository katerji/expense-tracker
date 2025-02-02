package user

import "errors"

type User struct {
	ID        uint32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Input struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

var (
	ErrFirstNameMissing    = errors.New("first name missing")
	ErrLastNameMissing     = errors.New("last name missing")
	ErrEmailMissing        = errors.New("email missing")
	ErrInvalidEmailMissing = errors.New("invalid email")
	ErrPassTooShort        = errors.New("password too short")
)

// TODO implement validator
func (i Input) validate() []error {
	errs := []error{}
	if i.FirstName == "" {
		errs = append(errs, ErrFirstNameMissing)
	}
	if i.LastName == "" {
		errs = append(errs, ErrLastNameMissing)
	}
	if i.Email == "" {
		errs = append(errs, ErrEmailMissing)
	}
	if i.Password == "" {
		errs = append(errs, ErrPassTooShort)
	}

	return errs
}
