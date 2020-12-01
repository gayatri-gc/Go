package owner

import(
	"fmt"
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct{
	DeleteOwner endpoint.Endpoint
	AddOwner endpoint.Endpoint
	GetOwner endpoint.Endpoint
	GetOwnersList endpoint.Endpoint
	UpdateOwnerDetails endpoint.Endpoint

	AddParkingSpot endpoint.Endpoint
	DeleteParkingSpot endpoint.Endpoint
	ParkVehicle endpoint.Endpoint
	UpdateParkingSpot endpoint.Endpoint
}

func MakeEndpoints(srv Service) Endpoints{
	return Endpoints{
		AddOwner: MakeAddOwnerEndpoint(srv),
		DeleteOwner: MakeDeleteOwnerEndpoint(srv),
		GetOwner: MakeGetOwnerEndpoint(srv),
		GetOwnersList: MakeGetOwnersListEndpoint(srv),
		UpdateOwnerDetails: MakeUpdateOwnerDetailsEndpoint(srv),

		AddParkingSpot: MakeAddParkingSpotEndpoint(srv),
		DeleteParkingSpot: MakeDeleteParkingSpotEndpoint(srv),
		ParkVehicle: MakeParkVehicleEndpoint(srv),
		UpdateParkingSpot: MakeUpdateParkingSpotEndpoint(srv),
	}
}

func MakeParkVehicleEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(ParkVehicleRequest)
		res, err := srv.ParkVehicle(ctx, req.RegNo, req.Color)
		return ParkVehicleResponse{Res: res}, err
	}
}

func MakeAddOwnerEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(AddOwnerRequest)
		owner := &OwnerDetails{
			req.OwnerID, 
			req.Name, 
			req.Email, 
			req.MobileNo,
		}
		fmt.Println("Inside MakeAddOwnerEndpoint ******* ", req)
		//err := srv.AddOwner(ctx, req.OwnerDetails)
		err := srv.AddOwner(ctx, owner)
		return AddOwnerResponse{Error: err}, err
	}
}

func MakeDeleteOwnerEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(DeleteOwnerRequest)
		err := srv.DeleteOwner(ctx, req.OwnerID)
		return DeleteOwnerResponse{Error: err}, err
	}
}

func MakeGetOwnerEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(GetOwnerRequest)
		result, err := srv.GetOwnerDetails(ctx, req.OwnerID)
		fmt.Println("Print owner data endpoint &&&&& ", result)
		return GetOwnerResponse{Owner:result, Error: err}, err
	}
}

func MakeGetOwnersListEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		result, err := srv.GetOwnersList(ctx)
		fmt.Println("print owners list endpoint******** ", result)
		return GetOwnersListResponse{List: result,Error: err}, err
	}
}

func MakeUpdateOwnerDetailsEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(UpdateOwnerDetailsRequest)
		ownerData := OwnerDetails{
			req.OwnerID, 
			req.Name, 
			req.Email, 
			req.MobileNo,
		}
		result, err := srv.UpdateOwnerDetails(ctx, ownerData)
		fmt.Println("print owners list endpoint******** ", result)
		return UpdateOwnerDetailsResponse{OwnerResp: result,Error: err}, err
	}
}

func MakeAddParkingSpotEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(AddParkingSpotRequest)
		spot := &ParkingSpot{
			req.SpotID,
			req.OwnerID, 
			req.Name, 
			req.GeoLongLat, 
			req.Address,
		}
	//	fmt.Println("Inside MakeAddOwnerEndpoint ******* ", owner)
		//err := srv.AddOwner(ctx, req.OwnerDetails)
		err := srv.AddParkingSpot(ctx, spot)
		return AddOwnerResponse{Error: err}, err
	}
}

func MakeDeleteParkingSpotEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(DeleteParkingSpotRequest)
		err := srv.DeleteParkingSpot(ctx, req.SpotID)
		return DeleteParkingSpotResponse{Error: err}, err
	}
}

func MakeUpdateParkingSpotEndpoint(srv Service) (endpoint.Endpoint){
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(UpdateParkingSpotRequest)
		ParkingSpotData := ParkingSpot{
			req.SpotID,
			req.OwnerID, 
			req.Name, 
			req.GeoLongLat, 
			req.Address,
		}
		result, err := srv.UpdateParkingSpot(ctx, ParkingSpotData)
		fmt.Println("print owners list endpoint******** ", result)
		return UpdateParkingSpotResponse{ParkingResp: result,Error: err}, err
	}
}