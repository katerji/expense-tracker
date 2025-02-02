package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo repo
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
