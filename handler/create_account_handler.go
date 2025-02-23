package handler

import (
	"encoding/json"
	"fmt"
	"github.com/katerji/expense-tracker/service/account"
	"github.com/katerji/expense-tracker/service/user"
	"net/http"
)

const CreateAccountRoute = "/account/create"

type CreateAccountRequest struct {
	Name string `json:"name"`
}

type CreateAccountResponse struct {
	JWTPair account.CustomJWTClaims `json:"jwt_pair"`
}

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		fmt.Println(err)
		return
	}
	ctx := r.Context()
	u := user.FromCtx(ctx)

	a, err := account.GetServiceInstance().GetUserAccount(ctx, u.ID)
	if err == nil || a != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("account exists"))
		return
	}

	a, err = account.GetServiceInstance().CreateAccount(ctx, account.CreateAccountInput{
		Name:   req.Name,
		UserID: u.ID,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong while creating account"))
		return
	}

	pair, err := account.GetServiceInstance().CreateJWTPair(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong while creating account"))
		return
	}

	w.WriteHeader(http.StatusOK)
	stringified, err := json.Marshal(pair)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong while creating account"))
		return
	}
	w.Write(stringified)
	return
}
