package fetcher

import "context"

type FetchItem struct {
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	TimeOfPurchase int64   `json:"string"`
	Description    string  `json:"description"`
	Merchant       string  `json:"merchant"`
	isValid        bool
}

func (f FetchItem) Valid() bool {
	return f.isValid
}

type Fetcher interface {
	Fetch(context.Context, []string) ([]FetchItem, bool)
}

func New() Fetcher {
	return service{}
}
