package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/katerji/expense-tracker/service/parser"
	"net/http"
)

const RecordTransactionRoute = "/transaction"

type RecordTransactionRequest struct {
	Message string `json:"message"`
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

	p := parser.NewParser()
	res, _ := p.Parse(context.Background(), req.Message)
	fmt.Println(res)

}
