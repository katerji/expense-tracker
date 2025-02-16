package parser

import (
	"context"
)

type transaction struct {
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	TimeOfPurchase int64   `json:"string"`
	Description    string  `json:"description"`
	Merchant       string  `json:"merchant"`
}

type TransactionDetailExtractor interface {
	Extract(context.Context, []string) ([]transaction, bool)
}

type detailExtractor struct{}

func (p detailExtractor) Extract(ctx context.Context, transactionMessages []string) ([]transaction, bool) {
	local := newLocalExtractor()
	if messages, ok := local.Extract(ctx, transactionMessages); ok {
		return messages, true
	}

	aiExtractor := newAIDetailExtractor()
	return aiExtractor.Extract(ctx, transactionMessages)
}

func NewDetailExtractor() TransactionDetailExtractor {
	return detailExtractor{}
}
