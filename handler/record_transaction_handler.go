package handler

import (
	"encoding/json"
	"fmt"
	"github.com/katerji/expense-tracker/service/parser"
	"net/http"
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

	transactionSplitter := parser.NewTransactionSplitter()
	transactionMessages, ok := transactionSplitter.Split(ctx, req.TransactionMessages)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		fmt.Println(err)
		return
	}
	p := parser.NewDetailExtractor()
	p.Extract(ctx, transactionMessages)
}
