package owner

import(
	//"github.com/jinzhu/gorm"
)

type OwnerDetails struct{
	//gorm.Model
	OwnerID uint  `gorm:"primary_key" json:"owner_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	MobileNo string `json:"mobile"`
	//ParkingSpot []ParkingSpot  `json:"parking_spot"`
}

type ParkingSpot struct{
	//gorm.Model  
	SpotID uint `gorm:"PRIMARY_KEY" json:"spot_id"`
	OwnerID uint`json:"owner_id"`
	Name string `json:"name"`
	GeoLongLat string `json:"geo_longlat"`
	Address string `json:"address"`
	//ParkingLotDetails *ParkingLotDetails `json:"parking_lot"`
}

type ParkingSpotInfo struct {
	Capacity  int `json:"capacity"`
	Slots     []Slot `json:"slots"`
	FreeSlots int `json:"free_slots"`
}

type VehicleDetails struct{
	RegNo string
	Color string 
}

type Slot struct {
	Occupied bool
}
