package handler

import (
	"encoding/json"
	"fmt"
	"github.com/katerji/expense-tracker/service/expense"
	"github.com/katerji/expense-tracker/service/parser"
	"github.com/katerji/expense-tracker/service/user"
	"net/http"
	"time"
)

const RecordTransactionRoute = "/transactions"

type RecordTransactionRequest struct {
	TransactionMessages string `json:"transactions"`
}

func RecordTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req RecordTransactionRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		fmt.Println(err)
		return
	}
	ctx := r.Context()
	account := user.AccountFromCtx(ctx)

	transactionSplitter := parser.NewTransactionSplitter()
	transactionMessages, ok := transactionSplitter.Split(ctx, req.TransactionMessages)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		fmt.Println(err)
		return
	}
	p := parser.NewDetailExtractor()
	transactions, ok := p.Extract(ctx, transactionMessages)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		fmt.Println(err)
		return
	}

	for _, transaction := range transactions {
		input := expense.RegisterExpenseInput{
			Amount:         transaction.Amount,
			Currency:       transaction.Currency,
			TimeOfPurchase: time.Unix(transaction.TimeOfPurchase, 0),
			Description:    "",
			MerchantName:   transaction.Merchant,
			MerchantType:   transaction.Merchant,
			AccountID:      account.ID,
		}
		expense.GetServiceInstance().RegisterExpense(ctx, input)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
	return
}
