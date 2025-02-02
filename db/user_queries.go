package db

import (
	"context"
	"github.com/katerji/expense-tracker/db/generated"
)

func InsertUser(ctx context.Context, params generated.InsertUserQueryParams) error {
	err := getInstance().InsertUserQuery(ctx, params)
	if err == nil {
		return nil
	}
	//TODO add logs
	//TODO parse error and return relevant errors to caller
	return err
}

func FetchUser(ctx context.Context, email string) (generated.FetchUserByEmailQueryRow, error) {
	res, err := getInstance().FetchUserByEmailQuery(ctx, email)
	if err == nil {
		return res, nil
	}
	//TODO add logs
	//TODO parse error and return relevant errors to caller
	return generated.FetchUserByEmailQueryRow{}, err
}

func FetchUserAccount(ctx context.Context, userID uint32) (generated.FetchUserAccountRow, error) {
	res, err := getInstance().FetchUserAccount(ctx, userID)
	if err == nil {
		return res, nil
	}
	//TODO add logs
	//TODO parse error and return relevant errors to caller
	return generated.FetchUserAccountRow{}, err
}

func InsertAccount(ctx context.Context, params generated.InsertAccountParams) error {
	return getInstance().InsertAccount(ctx, params)
}
