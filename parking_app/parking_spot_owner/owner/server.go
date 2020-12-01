package owner

import(
	"context"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler{
	fmt.Println("Inside NewHttpServer *** ")
	r := mux.NewRouter()
	r.Use(common)

	fmt.Println(" ************* ")

	r.Methods("POST").Path("/add").Handler(httptransport.NewServer(
		endpoints.AddOwner,
		decodeAddOwnerRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/add-parking-spot").Handler(httptransport.NewServer(
		endpoints.AddParkingSpot,
		decodeAddParkingSpotRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/delete-parking-spot").Handler(httptransport.NewServer(
		endpoints.DeleteParkingSpot,
		decodeDeleteParkingSpotRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/delete").Handler(httptransport.NewServer(
		endpoints.DeleteOwner,
		decodeDeleteOwnerRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/get").Handler(httptransport.NewServer(
		endpoints.GetOwner,
		decodeGetOwnerRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/update-owner").Handler(httptransport.NewServer(
		endpoints.UpdateOwnerDetails,
		decodeUpdateOwnerDetailsRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/list").Handler(httptransport.NewServer(
		endpoints.GetOwnersList,
		decodeGetOwnersListRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/update-parking-spot").Handler(httptransport.NewServer(
		endpoints.UpdateParkingSpot,
		decodeUpdateParkingSpotRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/park").Handler(httptransport.NewServer(
		endpoints.ParkVehicle,
		decodeParkVehicleRequest,
		encodeResponse,
	))
	return r
}


func common(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("content-Type", "application/json")
		handle.ServeHTTP(w, r)
	})
}
