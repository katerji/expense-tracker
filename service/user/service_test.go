package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestCreateUser(t *testing.T) {
	service := Service{}

	input := Input{
		FirstName: uuid.NewString(),
		LastName:  uuid.NewString(),
		Email:     fmt.Sprintf("%s@gmail.com", uuid.NewString()),
		Password:  uuid.NewString(),
	}

	u, errs := service.createUser(context.Background(), input)
	if len(errs) != 0 {
		for _, err := range errs {
			t.Error(err)
		}
		t.Fatal("")
	}

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
}
