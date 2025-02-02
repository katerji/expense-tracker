package user

import (
	"context"
	"errors"
	"github.com/katerji/expense-tracker/db"
	"github.com/katerji/expense-tracker/db/generated"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUnknown            = errors.New("unknown")
)

type repo struct{}

func (r repo) insertUser(ctx context.Context, userInput Input) (*User, error) {
	params := generated.InsertUserQueryParams{
		Email:     userInput.Email,
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		Password:  userInput.Password,
	}

	err := db.InsertUser(ctx, params)
	if err != nil {
		return nil, ErrEmailAlreadyExists
	}

	return r.fetchUser(ctx, userInput.Email)
}

func (r repo) fetchUser(ctx context.Context, email string) (*User, error) {
	res, err := db.FetchUser(ctx, email)
	if err != nil {
		return nil, ErrUnknown
	}

	return &User{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
	}, nil
}
