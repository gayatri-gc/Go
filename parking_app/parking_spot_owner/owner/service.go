package owner

import(
	"context"
)

type Service interface {
	AddOwner(ctx context.Context, owner *OwnerDetails) (err error)
	DeleteOwner(ctx context.Context, ownerID int) (err error)
	UpdateOwnerDetails(ctx context.Context, owner OwnerDetails) (ownerDetails OwnerDetails, err error)
	GetOwnerDetails(ctx context.Context, ownerID int) (OwnerDetails, error)
	GetOwnersList(ctx context.Context)([]OwnerDetails, error)

	AddParkingSpot(ctx context.Context, parkingSpot *ParkingSpot) (err error)
	DeleteParkingSpot(ctx context.Context, spotID uint) (err error)
	UpdateParkingSpot(ctx context.Context, parkingSpotData ParkingSpot)(parkingSpot ParkingSpot, err error)
	//ParkingSpotStatus(ctx context.Context, parkingSpot *ParkingSpot) (parkingLotDetails *ParkingLotDetails)
	ParkVehicle(ctx context.Context, regNo, color string) (string,error) 
	// VehicleStatus()
}