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
