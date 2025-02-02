package user

import (
	"context"
	"github.com/google/uuid"
	"testing"
)

func TestService_CreateAccount(t *testing.T) {
	t.Parallel()
	random, _ := uuid.NewUUID()

	input := CreateAccountInput{
		Name:   uuid.NewString(),
		UserID: random.ID(),
	}

	service := Service{}
	account, err := service.CreateAccount(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}

	if account.Name != input.Name {
		t.Errorf("expected Name %s, got %s instead", input.Name, account.Name)
	}
}
