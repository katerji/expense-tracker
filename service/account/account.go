package account

import "context"

type Account struct {
	ID   uint32
	Name string
}

type CreateAccountInput struct {
	Name   string
	UserID uint32
}

func (s *Service) GetUserAccount(ctx context.Context, userID uint32) (*Account, error) {
	return s.repo.fetchUserAccount(ctx, userID)
}

func (s *Service) GetAccountByID(ctx context.Context, id uint32) (*Account, error) {
	return s.repo.fetchAccountByID(ctx, id)
}

func (s *Service) CreateAccount(ctx context.Context, input CreateAccountInput) (*Account, error) {
	err := s.repo.insertAccount(ctx, input)
	if err != nil {
		return nil, err
	}

	return s.GetUserAccount(ctx, input.UserID)
}

func (a *Account) Ctx(ctx context.Context) context.Context {
	return context.WithValue(ctx, Account{}, a)
}

func FromCtx(ctx context.Context) Account {
	a, ok := ctx.Value(Account{}).(Account)
	if !ok {
		return Account{}
	}

	return a
}
