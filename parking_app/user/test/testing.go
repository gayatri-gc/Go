package test

import(
	"testing"
//	"../parking_app/user"
)

func Test_park(t *testing.T){
	p := user.NewParkingLot(10)
	p.ParkVehicle("MH-1232", "White")

}