package owner

import(
	"context"
	"errors"
	"fmt"
	//"log"
)

var cap = 10

type service struct{
	repo OwnerRepository
	park *ParkingSpotInfo
}

func NewService(userRepo OwnerRepository) Service{
	fmt.Println("Inside NewService ********")
	p := &ParkingSpotInfo{cap, make([]Slot, cap), cap}
	return &service{
		repo: userRepo,
		park: p,
	}
}

func init(){
	p := &ParkingSpotInfo{cap, make([]Slot, cap), cap}
	fmt.Println("Print parking lot ******* ", p)
}

func (p *ParkingSpotInfo)availableSlot() int {
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
	fmt.Println("Inside logic/ParkVehicle *********** ", regNo, color)
	//p := &ParkingLotDetails{cap, make([]Slot, cap), cap}
	
	//var p *ParkingLotDetails
	if serv.park == nil {
		fmt.Println("Parking lot does not exixts.")
		return "Not ok", errors.New("Parking lot does not exixts.")
	}
	slotNo := serv.park.availableSlot()
	if slotNo == -1 {
		fmt.Printf("Parking lot is full\n")
		return "Not ok", errors.New("Parking lot is full.")
	}                         
	serv.park.FreeSlots -= 1
	fmt.Println("Allocated slot number: ", slotNo+1, serv.park.FreeSlots)


	Vdata := &VehicleDetails{
		RegNo: regNo,
		Color: color,
	}
	err := serv.repo.ParkVehicle(ctx, Vdata)
	if err != nil{
		return "Failed to create entry", err
	}
	return "Vehicle parked successfully", nil

}

func (serv *service)AddOwner(ctx context.Context, owner *OwnerDetails) (err error){
	fmt.Println("Inside AddOwner ******* ", owner)
	err = serv.repo.AddOwner(ctx, owner)
	return err
}

func (serv *service)DeleteOwner(ctx context.Context, ownerID int) (err error){
	err = serv.repo.DeleteOwner(ctx, ownerID)
	return err
}

func (serv *service)GetOwnerDetails(ctx context.Context, ownerID int) (OwnerDetails, error){
	owner, err := serv.repo.GetOwnerDetails(ctx, ownerID)
	fmt.Println("Print owner data &&&&& ", owner)
	return owner, err
}

func (serv *service)GetOwnersList(ctx context.Context)([]OwnerDetails, error){
	owners, err := serv.repo.GetOwnersList(ctx)
	fmt.Println("print owners list logic******** ", owners)
	return owners, err
}

func (serv *service)UpdateOwnerDetails(ctx context.Context, owner OwnerDetails) (ownerDetails OwnerDetails, err error){
	ownerDetails, err = serv.repo.UpdateOwnerDetails(ctx, owner)
	if err != nil{
		return ownerDetails, err
	}
	return ownerDetails, nil
}


func (serv *service)AddParkingSpot(ctx context.Context, parkingSpot *ParkingSpot) (err error){
	err = serv.repo.AddParkingSpot(ctx, parkingSpot)
	if err != nil{
		return err
	}
	return nil
}

func (serv *service)DeleteParkingSpot(ctx context.Context, spotID uint) (err error){
	err = serv.repo.DeleteParkingSpot(ctx, spotID)
	return err
}

func (serv *service)UpdateParkingSpot(ctx context.Context, parkingSpotData ParkingSpot)(parkingSpot ParkingSpot, err error){
	parkingSpot, err = serv.repo.UpdateParkingSpot(ctx, parkingSpotData)
	fmt.Println("Inside UpdateParkingSpot ******** ", parkingSpot)
	return parkingSpot, err
}


// func ParkingSpotStatus(ctx context.Context, parkingSpot *ParkingSpot) (parkingLotDetails *ParkingLotDetails){

// }
