package communicator

import "context"

type Communicator interface {
	Get(context.Context, string) (map[string]any, bool)
}

func NewCommunicator() Communicator {
	return openAI{}
}
