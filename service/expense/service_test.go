package expense

import (
	"context"
	"github.com/google/uuid"
	"reflect"
	"testing"
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
