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

func FetchUserByEmail(ctx context.Context, email string) (generated.FetchUserByEmailQueryRow, error) {
	res, err := getInstance().FetchUserByEmailQuery(ctx, email)
	if err == nil {
		return res, nil
	}
	//TODO add logs
	//TODO parse error and return relevant errors to caller
	return generated.FetchUserByEmailQueryRow{}, err
}

func FetchUserByID(ctx context.Context, id uint32) (generated.FetchUserByIDRow, error) {
	res, err := getInstance().FetchUserByID(ctx, id)
	if err == nil {
		return res, nil
	}
	//TODO add logs
	//TODO parse error and return relevant errors to caller
	return generated.FetchUserByIDRow{}, err
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

func FetchAccountByID(ctx context.Context, id uint32) (generated.FetchAccountByIDRow, error) {
	res, err := getInstance().FetchAccountByID(ctx, id)
	if err == nil {
		return res, nil
	}
	//TODO add logs
	//TODO parse error and return relevant errors to caller
	return generated.FetchAccountByIDRow{}, err
}

func InsertAccount(ctx context.Context, params generated.InsertAccountParams) error {
	return getInstance().InsertAccount(ctx, params)
}
