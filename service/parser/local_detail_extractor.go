package parser

import (
	"context"
)

type localExtractor struct{}

func (l localExtractor) Extract(_ context.Context, _ []string) ([]transaction, bool) {
	return nil, false
}

func newLocalExtractor() TransactionDetailExtractor {
	return localExtractor{}
}
