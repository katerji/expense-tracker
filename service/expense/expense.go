package expense

import (
	"database/sql"
	"github.com/katerji/expense-tracker/db/generated"
	"time"
)

type Expense struct {
	ID             uint32
	Amount         float64
	Currency       string
	TimeOfPurchase time.Time
	Description    string
	MerchantID     uint32
	AccountID      uint32
}

type CreateInput struct {
	Amount         float64
	Currency       string
	TimeOfPurchase time.Time
	Description    string
	MerchantID     uint32
	AccountID      uint32
}

func (i CreateInput) queryParams() (generated.InsertExpenseQueryParams, bool) {
	return generated.InsertExpenseQueryParams{
		Amount:   i.Amount,
		Currency: i.Currency,
		TimeOfPurchase: sql.NullTime{
			Time:  i.TimeOfPurchase,
			Valid: !i.TimeOfPurchase.IsZero(),
		},
		Description: sql.NullString{
			String: i.Description,
			Valid:  i.Description != "",
		},
		MerchantID: i.MerchantID,
		AccountID:  i.AccountID,
	}, true
}

func (i CreateInput) expense(id uint32) *Expense {
	return &Expense{
		ID:             id,
		Amount:         i.Amount,
		Currency:       i.Currency,
		TimeOfPurchase: i.TimeOfPurchase,
		Description:    i.Description,
		MerchantID:     i.MerchantID,
		AccountID:      i.AccountID,
	}
}
