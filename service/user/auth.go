package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailNotFound   = errors.New("email not found")
	ErrInvalidPassword = errors.New("invalid password")
)

func (s *Service) Register(ctx context.Context, input Input) (*LoginResult, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}
	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		return nil, ErrUnknown
	}

	input.Password = hashedPassword

	_, err = s.repo.insertUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return s.Login(ctx, LoginInput{
		Email:    input.Email,
		Password: input.Password,
	})
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginResult struct {
	User    *User
	JWTPair *jwtPair
}

func (s *Service) Login(ctx context.Context, input LoginInput) (*LoginResult, error) {
	user, err := s.fetchUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, ErrEmailNotFound
	}

	if !validPassword(user.Password, input.Password) {
		return nil, ErrInvalidPassword
	}

	pair, err := s.createJWTPair(user)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		User:    user,
		JWTPair: pair,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
