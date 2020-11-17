package user

import (
	"context"
	"encoding/json"

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
)

func EncodeError(ctx context.Context, err error, w http.ResponseWriter){
	w.Header().Set("content-type:","application/json; charset:utf-8")
	w.WriteHeader(err2Code(err))
	json.NewEncoder(w).Encode(ErrorWrapper{Error: err.Error()})
}

func err2Code(err error) int{
	return http.StatusInternalServerError
}

type ErrorWrapper struct{
	Error string `json:"error"`
}

type Failure interface{
	Failed() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error){
	if f, ok := response.(Failure); !ok{
		w.Header().Set("content-type:","application/json; charset:utf-8")
		EncodeError(ctx, f.Failed(), w)
	}
	w.Header().Set("content-type:", "application/json; charset: utf-8")
	err = json.NewEncoder(w).Encode(response)
	return 
}

func decodeParkVehicleRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request ParkVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	return request, nil
}