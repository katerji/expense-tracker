package user

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

func (s *Service) CreateAccount(ctx context.Context, input CreateAccountInput) (*Account, error) {
	err := s.repo.insertAccount(ctx, input)
	if err != nil {
		return nil, err
	}

	return s.GetUserAccount(ctx, input.UserID)
}
