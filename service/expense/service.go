package expense

import "context"

type Service struct {
	repo repo
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
