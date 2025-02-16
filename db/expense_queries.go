package db

import (
	"context"
	"github.com/katerji/expense-tracker/db/generated"
)

func InsertExpense(ctx context.Context, input generated.InsertExpenseQueryParams) (uint32, bool) {
	res, err := getInstance().InsertExpenseQuery(ctx, input)
	if err != nil {
		return 0, false
	}

	return fromLastInsertIDtoUint32(res.LastInsertId())
}
