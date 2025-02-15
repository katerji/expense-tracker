package fetcher

import "context"

type dbFetcher struct{}

func (f dbFetcher) Fetch(_ context.Context, _ []string) ([]FetchItem, bool) {
	return nil, false
}

func newDBFetcher() Fetcher {
	return dbFetcher{}
}
