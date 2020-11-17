package service


import (
	"fmt"
	"log"

	"authentication/data"
)

// User encapsulate a user in the system.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

type RegPreChecker interface{
	UserExists(email string) bool
}

type RegPrecheck struct{}

var regCond RegPreChecker

func init(){
	regCond = RegPrecheck{}
}

func (reg RegPrecheck)UserExists(email string) bool{
	return data.UserExists(email)
}

// RegisterUser will register a User if only User has not been previously
// registered.
func RegisterUser(user User) error {
	// check if user is already registered
	found := regCond.UserExists(user.Email)
	if found {
		fmt.Println("Inside RegisterUser")
		return fmt.Errorf("email '%s' already registered", user.Email)
	}
	// carry business logic and Register the user in the system
	log.Println(user)
	return nil
}