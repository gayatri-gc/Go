package user

type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string	`json:"password"`
	Role string `json:"role"`
}

type ParkingLotDetails struct {
	Capacity  int
	Slots     []Slot
	FreeSlots int
}

type VehicleDetails struct{
	RegNo string
	Color string 
}

type LoginDetails struct {
	Token     string `json:"token"`
	//Verified  bool   `json:"isverified"`
	//Usertoken string `json:"usertoken"`
}

type Slot struct {
	Occupied bool
}
