package user

import(
	"fmt"
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct{
	ParkVehicle endpoint.Endpoint
}

func MakeEndpoints(srv Service) Endpoints{
	fmt.Println("Inside MakeEndpoints ********")
	return Endpoints{
		ParkVehicle: MakeParkVehicleEndpoint(srv),
	}
}

func MakeParkVehicleEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(ParkVehicleRequest)
		res, err := srv.ParkVehicle(ctx, req.RegNo, req.Color)
		return ParkVehicleResponse{Res: res}, err
	}
}

// type Endpoints struct{
// 	ParkVehicle endpoint.Endpoint
// }

// func MakeEndpoints(srv Service) (endpoint.Endpoint){
// 	return MakeParkVehicleEndpoint(srv)
// }

// func MakeParkVehicleEndpoint(srv Service) (endpoint.Endpoint){
// 	return func(ctx context.Context, request interface{})(interface{}, error){
// 		req := request.(ParkVehicleRequest)
// 		res, err := srv.ParkVehicle(ctx, req.RegNo, req.Color)
// 		return ParkVehicleResponse{Res: res}, err
// 	}
// }