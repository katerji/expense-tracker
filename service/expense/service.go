package expense

import (
	"context"
)

type Service struct {
	repo repo
}

var instance *Service

func GetServiceInstance() *Service {
	if instance == nil {
		instance = &Service{
			repo: repo{},
		}
	}

	return instance
}
func (s *Service) getOrCreateMerchant(ctx context.Context, merchantName, merchantTypeName string) (*merchant, bool) {
	mType, ok := s.repo.getOrInsertMerchantType(ctx, merchantTypeName)
	if !ok {
		return nil, false
	}

	return s.repo.getOrInsertMerchant(ctx, merchantName, mType)
}

func (s *Service) getMerchantByID(ctx context.Context, id uint32) (*merchant, bool) {
	return s.repo.getMerchantByID(ctx, id)
}

func (s *Service) RegisterExpense(ctx context.Context, input RegisterExpenseInput) (*Expense, bool) {
	merchant, ok := s.getOrCreateMerchant(ctx, input.MerchantName, input.MerchantType)
	if !ok {
		return nil, false
	}

	createInput := createExpenseInput{
		Amount:         input.Amount,
		Currency:       input.Currency,
		TimeOfPurchase: input.TimeOfPurchase,
		Description:    input.Description,
		MerchantID:     merchant.ID,
		AccountID:      input.AccountID,
	}

	return s.createExpense(ctx, createInput)
}

func (s *Service) createExpense(ctx context.Context, input createExpenseInput) (*Expense, bool) {
	return s.repo.insertExpense(ctx, input)
}
