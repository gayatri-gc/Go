package user

import(
	"context"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) (http.Handler){
	fmt.Println("Inside NewHttpServer *** ")
	r := mux.NewRouter()
	r.Use(common)
	r.Methods("GET").Path("/park").Handler(httptransport.NewServer(
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