package user

import(
	"context"
)

type UserRepository interface{
	ParkVehicle(ctx context.Context, Vdata VehicleDetails) error
}