package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/katerji/expense-tracker/service/ai"
	"net/url"
	"strings"
)

type aiDetailExtractor struct{}

func newAIDetailExtractor() TransactionDetailExtractor {
	return aiDetailExtractor{}
}

func (ade aiDetailExtractor) Extract(ctx context.Context, messages []string) ([]transaction, bool) {
	instance := ai.New()

	aiResult, ok := instance.Search(ctx, getDefaultAIMessage(messages))
	if !ok {
		return nil, false
	}

	return ade.parseAIResult(aiResult)

}

func getDefaultAIMessage(messages []string) string {
	messageSuffix := `Extract the following details from the given transaction messages that are separated by a new line: 
			- **amount**: The transaction amount (numerical). 
			- **currency**: The currency of the transaction (e.g., AED, EUR, USD). 
			- **merchant**: The name of the merchant or store where the transaction occurred. 
			- **merchant_type**: The type of business (e.g., restaurant, groceries, entertainment, utilities, household, etc.). 
			- **time_of_purchase**: The date and time of purchase in 'YYYY-MM-DD HH:MM:SS' format. 
			Please use the merchant name and transaction details to determine the 'merchant_type'. 
			If the merchant type is unclear from the transaction description, use the merchant name, location, and transaction context to infer the category. 
			If you still cannot determine the merchant type, set it to 'null'. 
			If any of the fields are missing or cannot be determined,
			**only then** search the web for additional details using the **currency** to help identify the country and merchant type. 
			If no information can be found, set the missing fields to 'null'.
			Return a json object with transactions as key, and value of array of JSON objects with these fields,
			Transactions: `

	messagesAsOneString := strings.Join(messages, "\n")
	messageToSend := fmt.Sprintf("%s%s", messageSuffix, messagesAsOneString)
	messageToSend = url.QueryEscape(messageToSend)

	return messageToSend
}

func (ade aiDetailExtractor) parseAIResult(aiResult map[string]any) ([]transaction, bool) {
	transactionsRaw, ok := aiResult["transactions"].([]any)
	if !ok {
		return nil, false
	}
	items := []transaction{}
	for _, transactionRaw := range transactionsRaw {
		transactionMap, ok := transactionRaw.(map[string]any)
		if !ok {
			continue
		}
		transactionAsString, err := json.Marshal(transactionMap)
		if err != nil {
			continue
		}
		var item transaction
		err = json.Unmarshal(transactionAsString, &item)
		if err != nil {
			continue
		}
		item.IsValid = true
		items = append(items, item)
	}

	return items, true
}
