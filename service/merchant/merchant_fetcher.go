package merchant

import "context"

type Fetcher interface {
	Fetch(context.Context, string) (map[string]any, bool)
}

type dbFetcher struct{}

func (f dbFetcher) Fetch(_ context.Context, _ string) (map[string]any, bool) {
	return nil, false
}

func NewFetcher() Fetcher {
	return dbFetcher{}
}
