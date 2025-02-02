package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestService_Register(t *testing.T) {
	service := Service{}

	input := Input{
		FirstName: uuid.NewString(),
		LastName:  uuid.NewString(),
		Email:     fmt.Sprintf("%s@gmail.com", uuid.NewString()),
		Password:  uuid.NewString(),
	}

	loginRes, err := service.Register(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}

	u := loginRes.User

	if u.FirstName != input.FirstName {
		t.Errorf("exepcted first name %s, got %s instead", input.FirstName, u.FirstName)
	}
	if u.LastName != input.LastName {
		t.Errorf("exepcted LastName %s, got %s instead", input.LastName, u.LastName)
	}
	if u.Email != input.Email {
		t.Errorf("exepcted Email %s, got %s instead", input.Email, u.Email)
	}
	if u.ID == 0 {
		t.Error("exepcted id to be set")
	}

	if loginRes.JWTPair.AccessToken == "" || loginRes.JWTPair.RefreshToken == "" {
		t.Fatal("expected AccessToken & Refresh Token to be set ")
	}
}
