package owner

import(
	"context"
)

type OwnerRepository interface {
	AddOwner(ctx context.Context, owner *OwnerDetails) (err error)
	DeleteOwner(ctx context.Context, ownerID int) (err error)
	GetOwnerDetails(ctx context.Context, ownerID int) (OwnerDetails, error)
	GetOwnersList(ctx context.Context)([]OwnerDetails, error)
	UpdateOwnerDetails(ctx context.Context, owner OwnerDetails) (ownerDetails OwnerDetails, err error)

	AddParkingSpot(ctx context.Context, parkingSpot *ParkingSpot) (err error)
	DeleteParkingSpot(ctx context.Context, spotID uint) (err error)
	UpdateParkingSpot(ctx context.Context, parkingSpotData ParkingSpot)(parkingSpot ParkingSpot, err error)

	//ParkingSpotStatus(ctx context.Context, parkingSpot *ParkingSpot) (parkingLotDetails *ParkingLotDetails)
	ParkVehicle(ctx context.Context, Vdata *VehicleDetails) error
	//VehicleStatus()
}