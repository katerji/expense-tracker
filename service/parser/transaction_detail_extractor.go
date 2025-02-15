package parser

import (
	"context"
	"github.com/katerji/expense-tracker/service/fetcher"
)

type TransactionDetailExtractor interface {
	Extract(context.Context, []string) ([]fetcher.FetchItem, bool)
}

type detailExtractor struct{}

func (p detailExtractor) Extract(ctx context.Context, transactionMessages []string) ([]fetcher.FetchItem, bool) {
	return fetcher.New().Fetch(ctx, transactionMessages)

}

func NewDetailExtractor() TransactionDetailExtractor {
	return detailExtractor{}
}
