package parser

import (
	"context"
	"errors"
	"github.com/katerji/expense-tracker/service/communicator"
	"github.com/katerji/expense-tracker/service/merchant"
	"time"
)

type parserResult struct {
	Amount         float64
	Currency       string
	TimeOfPurchase int64
	Description    string
	Merchant       string
}

var (
	ErrInvalidAmount   = errors.New("invalid amount")
	ErrInvalidCurrency = errors.New("invalid currency")
	ErrInvalidMerchant = errors.New("invalid merchant")
)

func parserResultFromMap(rawJson map[string]any) (*parserResult, error) {
	var amount float64
	var currency, description, merchantName string
	var timeOfPurchase int64

	if amountJSON, found := rawJson["amount"]; found {
		if amountJSONString, ok := amountJSON.(float64); ok {
			amount = amountJSONString
		} else {
			return nil, ErrInvalidAmount
		}
	}

	if currencyJSON, found := rawJson["currency"]; found {
		if currencyJSONString, ok := currencyJSON.(string); ok {
			currency = currencyJSONString
		} else {
			return nil, ErrInvalidCurrency
		}
	}
	if descriptionJSON, found := rawJson["description"]; found {
		if descriptionJSONString, ok := descriptionJSON.(string); ok {
			description = descriptionJSONString
		}
	}

	if timeOfPurchaseJSON, found := rawJson["time_of_purchase"]; found {
		if timeOfPurchaseJSONString, ok := timeOfPurchaseJSON.(string); ok {
			if timeOfPurchaseTime, err := time.Parse(time.DateTime, timeOfPurchaseJSONString); err != nil {
				timeOfPurchase = timeOfPurchaseTime.Unix()
			}
		}
	}

	if MerchantJSON, found := rawJson["merchant"]; found {
		if MerchantJSONString, ok := MerchantJSON.(string); ok {
			merchantName = MerchantJSONString
		} else {
			return nil, ErrInvalidMerchant
		}
	}

	return &parserResult{
		Amount:         amount,
		Currency:       currency,
		TimeOfPurchase: timeOfPurchase,
		Description:    description,
		Merchant:       merchantName,
	}, nil
}

type Parser interface {
	Parse(context.Context, string) (*parserResult, bool)
}

type genericParser struct{}

func (p genericParser) Parse(ctx context.Context, message string) (*parserResult, bool) {
	fetcher := merchant.NewFetcher()
	fetchPayload, ok := fetcher.Fetch(ctx, message)
	if ok {
		result, err := parserResultFromMap(fetchPayload)
		if err != nil {
			return result, true
		}
		// TODO add logs
	}

	comm := communicator.NewCommunicator()
	fetchPayload, ok = comm.Get(ctx, message)
	if ok {
		result, err := parserResultFromMap(fetchPayload)
		if err != nil {
			return result, true
		}
		// TODO add logs
	}

	return nil, false
}

func NewParser() Parser {
	return genericParser{}
}
