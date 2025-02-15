package ai

import "context"

type AI interface {
	Search(ctx context.Context, message string) (map[string]any, bool)
}

func New() AI {
	return &openAI{}
}
