package db

import (
	"context"
	"github.com/katerji/expense-tracker/db/generated"
)

func FetchMerchantByName(ctx context.Context, name string) (generated.FetchMerchantByNameQueryRow, error) {
	return getInstance().FetchMerchantByNameQuery(ctx, name)
}

func FetchMerchantType(ctx context.Context, typeName string) (*generated.MerchantType, bool) {
	res, err := getInstance().FetchMerchantTypeQuery(ctx, typeName)
	if err != nil {
		return nil, false
	}

	return &res, true
}

func InsertMerchantType(ctx context.Context, typeName string) (uint32, bool) {
	res, err := getInstance().InsertMerchantTypeQuery(ctx, typeName)
	if err != nil {
		return 0, false
	}

	return fromLastInsertIDtoUint32(res.LastInsertId())
}

func InsertMerchant(ctx context.Context, input generated.InsertMerchantQueryParams) (uint32, bool) {
	res, err := getInstance().InsertMerchantQuery(ctx, input)
	if err != nil {
		return 0, false
	}

	return fromLastInsertIDtoUint32(res.LastInsertId())
}
