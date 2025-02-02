package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

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

func (s *Service) fetchUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.repo.fetchUser(ctx, email)
}

func (s *Service) createUser(ctx context.Context, userInput Input) (*User, []error) {
	if errs := userInput.validate(); len(errs) > 0 {
		return nil, errs
	}
	hashedPassword, err := hashPassword(userInput.Password)
	if err != nil {
		return nil, []error{ErrUnknown}
	}

	userInput.Password = hashedPassword

	user, err := s.repo.insertUser(ctx, userInput)
	if err != nil {
		return nil, []error{err}
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
