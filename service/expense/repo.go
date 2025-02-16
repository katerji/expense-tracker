package expense

import (
	"context"
	"github.com/katerji/expense-tracker/db"
)

type repo struct{}

func (r repo) insert(ctx context.Context, input CreateInput) (*Expense, bool) {
	queryParams, ok := input.queryParams()
	if !ok {
		return nil, false
	}
	insertID, ok := db.InsertExpense(ctx, queryParams)
	if !ok {
		return nil, false
	}

	return input.expense(insertID), true
}

func (r repo) fetchMerchantByName(ctx context.Context, merchantName string) (*merchant, bool) {
	queryRes, err := db.FetchMerchantByName(ctx, merchantName)
	if err != nil {
		return nil, false
	}

	return &merchant{
		ID:     queryRes.ID,
		Name:   queryRes.Name,
		TypeID: queryRes.TypeID,
	}, true
}
