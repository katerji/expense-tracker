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

	return r.fetchUserByEmail(ctx, userInput.Email)
}

func (r repo) fetchUserByEmail(ctx context.Context, email string) (*User, error) {
	res, err := db.FetchUserByEmail(ctx, email)
	if err != nil {
		return nil, ErrUnknown
	}

	return &User{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
		Password:  res.Password,
	}, nil
}

func (r repo) fetchUserByID(ctx context.Context, id uint32) (*User, error) {
	res, err := db.FetchUserByID(ctx, id)
	if err != nil {
		return nil, ErrUnknown
	}

	return &User{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
		Password:  res.Password,
	}, nil
}

func (r repo) fetchUserAccount(ctx context.Context, userID uint32) (*Account, error) {
	res, err := db.FetchUserAccount(ctx, userID)
	if err != nil {
		return nil, ErrUnknown
	}

	return &Account{
		ID:   res.ID,
		Name: res.Name,
	}, nil
}
func (r repo) fetchAccountByID(ctx context.Context, userID uint32) (*Account, error) {
	res, err := db.FetchUserAccount(ctx, userID)
	if err != nil {
		return nil, ErrUnknown
	}

	return &Account{
		ID:   res.ID,
		Name: res.Name,
	}, nil
}

func (r repo) insertAccount(ctx context.Context, input CreateAccountInput) error {
	params := generated.InsertAccountParams(input)
	err := db.InsertAccount(ctx, params)
	if err == nil {
		return nil
	}

	return ErrUnknown
}
