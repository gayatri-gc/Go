package user

import(
	"context"
)

type UserRepository interface{
	Register(ctx context.Context, user *User) (usr *User, err error)
	Login(ctx context.Context, user *User) (usr User, err error)
	ParkVehicle(ctx context.Context, Vdata VehicleDetails) error
}