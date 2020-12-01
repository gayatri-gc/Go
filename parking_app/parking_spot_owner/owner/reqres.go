package owner

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
	
	AddOwnerRequest struct{
		//OwnerDetails *OwnerDetails `json:"owner"`
		OwnerID uint  `json:"owner_id"`
		Name string `json:"name"`
		Email string `json:"email"`
		MobileNo string `json:"mobile"`
	}

	AddOwnerResponse struct {
		Error error `json:"error"`
	}

	DeleteOwnerRequest struct{
		OwnerID int `json:"owner_id"`
	}

	DeleteOwnerResponse struct{
		Error error `json:"error"`
	}

	GetOwnerRequest struct{
		OwnerID int `json:"owner_id"`
	}

	GetOwnerResponse struct{
		Owner OwnerDetails `json:"owner"`
		// ID int  `json:"id"`
		// Name string `json:"name"`
		// Email string `json:"email"`
		// MobileNo string `json:"mobile"`
		Error error `json:"error"`
	}

	GetOwnersListRequest struct{}

	GetOwnersListResponse struct{
		List []OwnerDetails `json:"list"`
		Error error `json:"error"`
	}

	UpdateOwnerDetailsRequest struct{
		//Owner OwnerDetails `json:"owner"`
		OwnerID uint  `json:"owner_id"`
		Name string `json:"name"`
		Email string `json:"email"`
		MobileNo string `json:"mobile"`
	}

	UpdateOwnerDetailsResponse struct{
		OwnerResp OwnerDetails `json:"owner"`
		Error error `json:"error"`
	}

	AddParkingSpotRequest struct{
		//ParkingSpot *ParkingSpot `json:"parking_spot"`
		SpotID uint `json:"spot_id"`
		OwnerID uint  `json:"owner_id"`
		Name string `json:"name"`
		GeoLongLat string `json:"geo_longlat"`
		Address string `json:"address"`
	}

	AddParkingSpotResponse struct {
		Error error `json:"error"`
	}

	DeleteParkingSpotRequest struct{
		SpotID uint `json:"spot_id"`
	}

	DeleteParkingSpotResponse struct{
		Error error `json:"error"`
	}

	UpdateParkingSpotRequest struct{
		//Owner OwnerDetails `json:"owner"`
		SpotID uint `json:"spot_id"`
		OwnerID uint  `json:"owner_id"`
		Name string `json:"name"`
		GeoLongLat string `json:"geo_longlat"`
		Address string `json:"address"`
	}

	UpdateParkingSpotResponse struct{
		ParkingResp ParkingSpot `json:"parking_spot"`
		Error error `json:"error"`
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
	fmt.Println("Inside encodeResponse ********* ", response)
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

func decodeParkVehicleRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request ParkVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("Inside reqres/decodeParkVehicleRequest *********** ", request)
	return request, nil
}

func decodeAddOwnerRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request AddOwnerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("Inside reqres/decodeAddOwnerRequest *********** ", request)
	return request, nil
}

func decodeDeleteOwnerRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request DeleteOwnerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("Inside reqres/decodeAddOwnerRequest *********** ", request)
	return request, nil
}

func decodeGetOwnerRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request GetOwnerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("Inside reqres/decodeGetOwnerRequest *********** ", request)
	return request, nil
}

func decodeGetOwnersListRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request GetOwnersListRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("Inside reqres/decodeGetOwnerRequest *********** ", request)
	return request, nil
}

func decodeUpdateOwnerDetailsRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request UpdateOwnerDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("Inside reqres/decodeUpdateOwnerDetailsRequest *********** ", request)
	return request, nil
}

func decodeAddParkingSpotRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request AddParkingSpotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("Inside reqres/decodeAddOwnerRequest *********** ", request)
	return request, nil
}

func decodeDeleteParkingSpotRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request DeleteParkingSpotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("Inside reqres/decodeAddOwnerRequest *********** ", request)
	return request, nil
}

func decodeUpdateParkingSpotRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request UpdateParkingSpotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	fmt.Println("Inside reqres/decodeUpdateOwnerDetailsRequest *********** ", request)
	return request, nil
}