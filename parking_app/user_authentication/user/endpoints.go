package user

import(
	"fmt"
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct{
	Register 	endpoint.Endpoint
	Login       endpoint.Endpoint
	ParkVehicle endpoint.Endpoint
}

func MakeEndpoints(srv Service) Endpoints{
	return Endpoints{
		Register: makeRegisterEndpoint(srv),
		Login:        makeLoginEndpoint(srv),
		ParkVehicle: MakeParkVehicleEndpoint(srv),
	}
}

func makeRegisterEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		user := &User{req.Name, req.Email,req.Password, req.Role}
		resp, err := srv.Register(ctx, user)
		if resp == nil{
			return "no record found", nil
		}
		reg := &RegisterResponse{
			Name: resp.Name,
			Email: resp.Email,
			Password: resp.Password,
			Error:       err,
		}
		return reg, nil
	}
}

func makeLoginEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		user := &User{req.Name, req.Email,req.Password, req.Role}
		resp, err := srv.Login(ctx, user)
		if resp == nil{
			return "no record found", nil
		}
		fmt.Println("Print resp from service ****** ", resp)
		return LoginResponse{
			LoginDetails: resp,
			Error:       err,
		}, nil
	}
}

func MakeParkVehicleEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(ParkVehicleRequest)
		res, err := srv.ParkVehicle(ctx, req.RegNo, req.Color)
		return ParkVehicleResponse{Res: res}, err
	}
}
