package expense

import (
	"context"
	"github.com/katerji/expense-tracker/db"
	"github.com/katerji/expense-tracker/db/generated"
)

type repo struct{}

func (r repo) insertExpense(ctx context.Context, input createExpenseInput) (*Expense, bool) {
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

func (r repo) insertMerchant(ctx context.Context, input createMerchantInput) (uint32, bool) {
	return db.InsertMerchant(ctx, generated.InsertMerchantQueryParams(input))
}

func (r repo) fetchMerchantByName(ctx context.Context, merchantName string) (*merchant, bool) {
	queryRes, ok := db.FetchMerchantByName(ctx, merchantName)
	if !ok {
		return nil, false
	}

	return &merchant{
		ID:   queryRes.ID,
		Name: queryRes.Name,
		merchantType: &merchantType{
			ID:   queryRes.TypeID,
			Type: queryRes.MerchantType,
		},
	}, true
}

func (r repo) getOrInsertMerchant(ctx context.Context, name string, mType *merchantType) (*merchant, bool) {
	m, ok := r.fetchMerchantByName(ctx, name)
	if ok {
		return m, true
	}

	merchantID, ok := r.insertMerchant(ctx, createMerchantInput{
		Name:   name,
		TypeID: mType.ID,
	})
	if !ok {
		return nil, false
	}

	return &merchant{
		ID:           merchantID,
		Name:         name,
		merchantType: mType,
	}, true
}

func (r repo) insertMerchantType(ctx context.Context, typeName string) (*merchantType, bool) {
	id, ok := db.InsertMerchantType(ctx, typeName)
	if !ok {
		return nil, false
	}

	return &merchantType{
		ID:   id,
		Type: typeName,
	}, true
}

func (r repo) fetchMerchantType(ctx context.Context, typeName string) (*merchantType, bool) {
	res, ok := db.FetchMerchantType(ctx, typeName)
	if !ok {
		return nil, false
	}

	return &merchantType{
		ID:   res.ID,
		Type: res.Type,
	}, true
}

func (r repo) getOrInsertMerchantType(ctx context.Context, typeName string) (*merchantType, bool) {
	mType, ok := r.fetchMerchantType(ctx, typeName)
	if ok {
		return mType, ok
	}

	return r.insertMerchantType(ctx, typeName)
}

func (r repo) getMerchantByID(ctx context.Context, id uint32) (*merchant, bool) {
	res, ok := db.FetchMerchantByID(ctx, id)
	if !ok {
		return nil, false
	}

	return &merchant{
		ID:   res.ID,
		Name: res.Name,
		merchantType: &merchantType{
			ID:   res.TypeID,
			Type: res.MerchantType,
		},
	}, true
}
