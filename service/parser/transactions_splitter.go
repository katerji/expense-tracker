package parser

import (
	"context"
	"fmt"
	"github.com/katerji/expense-tracker/service/ai"
	"net/url"
)

type TransactionsSplitter interface {
	Split(ctx context.Context, transactions string) ([]string, bool)
}

func NewTransactionSplitter() TransactionsSplitter {
	return splitter{}
}

type splitter struct{}

func (s splitter) Split(ctx context.Context, transactions string) ([]string, bool) {
	aiInstance := ai.New()
	aiResult, ok := aiInstance.Search(ctx, getAIMessage(transactions))
	if !ok {
		return nil, false
	}

	return s.parseAIResult(aiResult)
}

func (s splitter) parseAIResult(aiResult map[string]any) ([]string, bool) {
	transactionsRaw, ok := aiResult["transactions"].([]any)
	if !ok {
		return nil, false
	}

	transactionsList := make([]string, 0, len(transactionsRaw)) // Preallocate capacity
	for _, t := range transactionsRaw {
		if transaction, ok := t.(string); ok {
			transactionsList = append(transactionsList, transaction)
		}
	}

	return transactionsList, true
}

func getAIMessage(transactions string) string {
	messageSuffix := "Extract individual transactions from the given text and return them as a JSON-formatted array of strings," +
		" ensuring each transaction is a standalone string containing only relevant details such as" +
		" amount, currency, date, time, and merchant name, while ignoring unrelated text like available credit limits. transactionMessages: "
	messageToSend := fmt.Sprintf("%s%s", messageSuffix, transactions)
	messageToSend = url.QueryEscape(messageToSend)

	return messageToSend
}
