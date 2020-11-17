package user

import(
	"context"
	"errors"
	"fmt"
)

const cap = 10

type service struct{
	repo UserRepository
}

func NewService(userRepo UserRepository) Service{
	fmt.Println("Inside NewService ********")
	return &service{
		repo: userRepo,
	}
}


func (p *ParkingLotDetails)availableSlot() int {
	if p.FreeSlots > 0 {
		n := p.Capacity
		for i := 0; i < n; i++ {
			if !p.Slots[i].Occupied {
				return i
			}
		}
	}
	return -1
}


func (serv *service)ParkVehicle(ctx context.Context, regNo string, color string) (string,error){
	p := &ParkingLotDetails{cap, make([]Slot, cap), cap}
	if p == nil {
		fmt.Println("Parking lot does not exixts.")
		return "Not ok", errors.New("Parking lot does not exixts.")
	}
	slotNo := p.availableSlot()
	if slotNo == -1 {
		fmt.Printf("Parking lot is full\n")
		return "Not ok", errors.New("Parking lot is full.")
	}
	fmt.Printf("Allocated slot number: %d\n", slotNo+1)


	Vdata := VehicleDetails{
		RegNo: regNo,
		Color: color,
	}
	err := serv.repo.ParkVehicle(ctx, Vdata)
	if err != nil{
		return "Failed to create entry", err
	}
	return "Vehicle parked successfully", nil

}
