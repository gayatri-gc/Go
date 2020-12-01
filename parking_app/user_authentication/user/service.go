package user

import(
	"context"
)

type Service interface{
	Register(ctx context.Context, user *User) (usr *User, err error)
	Login(ctx context.Context, user *User) (loginDetails *LoginDetails, err error)

	ParkVehicle(ctx context.Context, regNo, color string) (string,error) 
	// GetVehicleBySlot(ctx context.Context, slotNo int) (*VehicleDetails, error)
	// GetVehicleByColor(ctx context.Context, color string) (*VehicleDetails, error)
	// GetVehicleByRegNo(ctx context.Context, regNo int) (*VehicleDetails, error)
	// FreeSlots(ctx context.Context, parkingLot string) (int, error)
	// Leave(ctx context.Context, slotNo int) (string)
	// ParkingLotstatus(ctx context.Context, parkingLot string) (*ParkingLotDetails, error)
}