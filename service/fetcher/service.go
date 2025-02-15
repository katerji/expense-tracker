package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/katerji/expense-tracker/service/ai"
	"net/url"
	"strings"
)

type service struct{}

func (s service) Fetch(ctx context.Context, messages []string) ([]FetchItem, bool) {
	return s.fetchFromAI(ctx, messages)

}

func (s service) fetchFromDB(ctx context.Context, messages []string) ([]FetchItem, bool) {
	db := newDBFetcher()

	return db.Fetch(ctx, messages)
}

func (s service) fetchFromAI(ctx context.Context, messages []string) ([]FetchItem, bool) {
	instance := ai.New()

	aiResult, ok := instance.Search(ctx, getDefaultAIMessage(messages))
	if !ok {
		return nil, false
	}

	return s.parseAIResult(aiResult)

}

func getDefaultAIMessageV1(messages []string) string {
	messageSuffix := "Extract the amount, currency, merchant, merchant_type (e.g., restaurant, groceries, entertainment, utilities, household, etc.), " +
		"and time_of_purchase (in YYYY-MM-DD HH:MM:SS format) from the following transaction messages that are separated by a new line. " +
		"Return a json object with transactions as key, and value of array of JSON objects with these fields.  " +
		"If any field is missing, set it to null. transactionMessages: "
	messagesAsOneString := strings.Join(messages, "\n")
	messageToSend := fmt.Sprintf("%s%s", messageSuffix, messagesAsOneString)
	messageToSend = url.QueryEscape(messageToSend)

	return messageToSend
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

func (s service) parseAIResult(aiResult map[string]any) ([]FetchItem, bool) {
	transactionsRaw, ok := aiResult["transactions"].([]any)
	if !ok {
		return nil, false
	}
	items := []FetchItem{}
	for _, transaction := range transactionsRaw {
		transactionMap, ok := transaction.(map[string]any)
		if !ok {
			continue
		}
		transactionAsString, err := json.Marshal(transactionMap)
		if err != nil {
			continue
		}
		var item FetchItem
		err = json.Unmarshal(transactionAsString, &item)
		if err != nil {
			continue
		}
		item.isValid = true
		items = append(items, item)
	}

	return items, true
}
