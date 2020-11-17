package test

import (
	"fmt"
	"testing"
	"authentication/service"
)

var UserExistMock func (email string) bool

type preCheck struct{}

func (pre preCheck) UserExists(email string)bool{
	return UserExistMock(email)
}

func TestCheckUserExist(t *testing.T) {
	fmt.Println("inside test ********* ")
	user := service.User{
		Name:     "Gayatri Chougale",
		Email:    "gayatri61995@gmail.com",
		UserName: "Gayatri",
	}

	//regPreCond := preCheck{}

	UserExistMock = func(email string) bool{
		return false
	}

	err := service.RegisterUser(user)
	if err != nil {
		t.Fatal(err)
	}
	
	UserExistMock = func(email string) bool{
		return true
	}

	err = service.RegisterUser(user)

	if err == nil {
		t.Error("Expected Register User to throw and error got nil")
	}
}