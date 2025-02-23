package expense

import (
	"context"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

func TestGetOrCreateMerchant(t *testing.T) {
	expectedMerchantName := uuid.NewString()
	expectedMerchantType := uuid.NewString()
	ctx := context.Background()

	s := Service{}
	merchant, ok := s.getOrCreateMerchant(ctx, expectedMerchantName, expectedMerchantType)
	if !ok {
		t.Fatalf("failed to get or create merchant")
	}

	if merchant == nil {
		t.Fatalf("failed to get or create merchant")
	}

	insertedMerchant, ok := s.getMerchantByID(ctx, merchant.ID)
	if !ok {
		t.Fatalf("failed to fetch inserted merchant with ID %d", merchant.ID)
	}

	if !reflect.DeepEqual(insertedMerchant, merchant) {
		t.Errorf("expected merchant %v, got %v", merchant, insertedMerchant)
	}
}

func TestRegisterExpense(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	registerExpenseInput := RegisterExpenseInput{
		Amount:         7,
		Currency:       "AED",
		TimeOfPurchase: time.Now(),
		Description:    uuid.NewString(),
		MerchantName:   uuid.NewString(),
		MerchantType:   uuid.NewString(),
		AccountID:      1,
	}
	service := Service{}

	expense, ok := service.RegisterExpense(ctx, registerExpenseInput)
	if !ok {
		t.Fatalf("failed to register expense")
	}

	if expense.Amount != registerExpenseInput.Amount {
		t.Errorf("expected amount %v, got %v instead", registerExpenseInput.Amount, expense.Amount)
	}

	if expense.Currency != registerExpenseInput.Currency {
		t.Errorf("expected currency %s, got %s instead", registerExpenseInput.Currency, expense.Currency)
	}

	if !expense.TimeOfPurchase.Equal(registerExpenseInput.TimeOfPurchase) {
		t.Errorf("expected time of purchase to be %d, got %d instead", registerExpenseInput.TimeOfPurchase.Unix(), expense.TimeOfPurchase.Unix())
	}

	if expense.Description != registerExpenseInput.Description {
		t.Errorf("expected Description %s, got %s instead", registerExpenseInput.Description, expense.Description)
	}

	if expense.MerchantID == 0 {
		t.Errorf("expected merchant id to be inserted correctly")
	}

	if expense.AccountID == 0 {
		t.Errorf("expected account id to be inserted correctly")
	}
}
