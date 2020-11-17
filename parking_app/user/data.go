package user

type ParkingLotDetails struct {
	Capacity  int
	Slots     []Slot
	FreeSlots int
}

type VehicleDetails struct{
	RegNo string
	Color string 
}

type Slot struct {
	Occupied bool
}
