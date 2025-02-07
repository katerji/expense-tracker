package parser

import (
	"context"
	"testing"
	"time"
)

const (
	successAmount      = 1.0
	successCurrency    = "AED"
	successDescription = "successful transaction"
	successMerchant    = "successful merchant"
)

var successTimeOfPurchase = time.Now().Unix()

type parserMockSuccess struct{}

func (p parserMockSuccess) Parse(_ context.Context, _ string) (*parserResult, bool) {
	return &parserResult{
		Amount:         successAmount,
		Currency:       successCurrency,
		TimeOfPurchase: successTimeOfPurchase,
		Description:    successDescription,
		Merchant:       successMerchant,
	}, true
}

func TestParseSuccess(t *testing.T) {
	parser := parserMockSuccess{}

	res, ok := parser.Parse(context.Background(), "dummy")
	if !ok {
		t.Fatalf("expected parser to succeed, it failed instaed")
	}

	if res.Currency != successCurrency {
		t.Errorf("expected currency %s, got %s instead", successCurrency, res.Currency)
	}

	if res.Currency != successCurrency {
		t.Errorf("expected currency %s, got %s instead", successCurrency, res.Currency)
	}

	if res.TimeOfPurchase != successTimeOfPurchase {
		t.Errorf("expected time of purchase %v, got %v instead", successTimeOfPurchase, res.TimeOfPurchase)
	}

	if res.Description != successDescription {
		t.Errorf("expected description %s, got %s instead", successDescription, res.Description)
	}

	if res.Merchant != successMerchant {
		t.Errorf("expected merchant %s, got %s instead", successMerchant, res.Merchant)
	}
}
