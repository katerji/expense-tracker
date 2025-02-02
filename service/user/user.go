package user

import (
	"context"
	"errors"
)

type User struct {
	ID        uint32
	FirstName string
	LastName  string
	Email     string
	Password  string
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
func (i Input) validate() error {
	if i.FirstName == "" {
		return ErrFirstNameMissing
	}
	if i.LastName == "" {
		return ErrLastNameMissing
	}
	if i.Email == "" {
		return ErrEmailMissing
	}
	if i.Password == "" {
		return ErrPassTooShort
	}

	return nil
}

func (s *Service) fetchUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.repo.fetchUser(ctx, email)
}
