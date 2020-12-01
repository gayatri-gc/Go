package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
)

type (
	ParkVehicleRequest struct {
		RegNo string  `json:"regno"`
		Color string  `json:"color"`
	}

	ParkVehicleResponse struct {
		Res string	`json:"res"`
		Error error      `json:"error"`
	}	

	RegisterRequest struct{
		Name string	`json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
		Role string `json:"role"`
		//User *User `json:"user"`
	}

	RegisterResponse struct{
		Name string	`json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
		//User *User `json:"user"`
		Error       error       `json:"error"`
	}
	
	LoginRequest struct {
		Name string	`json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
		Role string `json:"role"`
		//User *User `json:"user"`
		//Email
	}

	//LoginResponse struct
	LoginResponse struct {
		LoginDetails *LoginDetails `json:"data"`
		Error       error       `json:"error"`
	}
)

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

func err2code(err error) int{
	return http.StatusInternalServerError
}

type errorWrapper struct{
	Error string `json:"error"`
}

type Failure interface{
	Failed() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

func decodeRegisterRequest(ctx context.Context, r *http.Request)(interface{}, error){
	fmt.Println("Inside decodeRegisterRequest *********** ")
	var request RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("get request data  *********** ", request)
	return request, nil
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("Inside decodeLoginRequest *********** ")
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}

	return request, nil
	// if (email == "" || password == "") {
	// 	ErrEmpty := errors.New("Provide name and Password")
	// 	return nil, ErrEmpty
	// }
	// req.User.Email = username
	// req.User.Password = password
	// return req, nil
}

func decodeParkVehicleRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request ParkVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	return request, nil
}