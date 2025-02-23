package account

import (
	"context"
	"errors"
	"github.com/katerji/expense-tracker/db"
	"github.com/katerji/expense-tracker/db/generated"
)

var ErrUnknown = errors.New("unknown")

type repo struct{}

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
