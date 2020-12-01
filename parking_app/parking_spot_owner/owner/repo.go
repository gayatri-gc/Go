package owner

import 
(
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
)

const ExpectedMinTime = 120

type OwnerRepo struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) OwnerRepository{
	fmt.Println("Inside NewRepository ********", db)
	return &OwnerRepo{
		db: db,
	}
} 

func(repo *OwnerRepo)ParkVehicle(ctx context.Context, Vdata *VehicleDetails) error{
	fmt.Println("Inside CreateUser ********* ", Vdata)
	repo.db.AutoMigrate(&VehicleDetails{})
	_ = repo.db.Create(&Vdata)
	return nil
}

func (repo *OwnerRepo)AddOwner(ctx context.Context, owner *OwnerDetails) (err error){
	fmt.Println("inside repo/AddOwner ******* ", owner)
	_ = repo.db.Create(&owner)
	return nil
}

func (repo *OwnerRepo)DeleteOwner(ctx context.Context, ownerID int) (err error){
	resp := repo.db.Delete(&OwnerDetails{}, ownerID)
	if resp.Error != nil{
		return resp.Error
	}
	return nil
}

func (repo *OwnerRepo)GetOwnerDetails(ctx context.Context, ownerID int) (OwnerDetails, error){
// 	fmt.Println("Inside GetOwnerDetails ******* ", ownerID)
// 	var ownerD OwnerDetails
// 	rows, err := repo.db.Where("owner_id = ?", ownerID).Find(&ownerD).Rows()
// 	ownerData := &OwnerDetails{}
// //	for rows.Next() {
//         //ownerItem := OwnerDetails{}
//        // item := Item{}
//         err = rows.Scan(&ownerData.OwnerID, &ownerData.Name, &ownerData.Email, &ownerData.MobileNo)
//         if err != nil {
//            // log.Panic(err)
//         }
//     //}
// 	// if result.Error != nil{
// 	// 	return owner, result.Error
// 	// }
// 	// //result.RowsAfftected
// 	// fmt.Println("Print owner from repo ******* ", result)
// 	// return ownerD, nil
// 	fmt.Println("Print owner data ***** ", ownerData)

// 	DB := repo.db.Debug().Model(&OwnerDetails{}).Where("owner_id = ?", ownerID).Find()
// 	return owner, DB.Error


	ownerdb := OwnerDetails{}
	data := repo.db.Debug().
		Where("owner_id=?", ownerID).First(&ownerdb)
	fmt.Println("&&&&&&&&&&&&&&&&&", ownerdb)	
	return ownerdb, data.Error
}

func (repo *OwnerRepo)GetOwnersList(ctx context.Context)([]OwnerDetails, error){
	//var ownerList []OwnerDetails
	rows, err := repo.db.Find(&OwnerDetails{}).Rows()

	owners := make([]OwnerDetails, 0)

    for rows.Next() {
        owner := OwnerDetails{}
        // orderItem := OrderItem{}

        // item := Item{}
        err = rows.Scan(&owner.OwnerID, &owner.Name, &owner.Email, &owner.MobileNo)
        if err != nil {
           fmt.Println(err)
        }

        owners = append(owners, owner)
    }
	fmt.Println("print owners list repo ******** ", owners)
	return owners, err
}

func (repo *OwnerRepo)UpdateOwnerDetails(ctx context.Context, owner OwnerDetails) (ownerDetails OwnerDetails, err error){
	fmt.Println("inside UpdateOwnerDetails ***** ", owner, owner.OwnerID, owner.Email)
	if err := repo.db.Model(&ownerDetails).Where("owner_id = ?", owner.OwnerID).Update(&owner).Error; err != nil {
		fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%")
		if gorm.IsRecordNotFoundError(err){
			fmt.Println("RecordNotFound")
			repo.db.Create(&ownerDetails)  // create new record from newUser
		}
	}
	fmt.Println("Print OwnerDetails ***********", ownerDetails )
	return ownerDetails, err
}


func (repo *OwnerRepo)AddParkingSpot(ctx context.Context, parkingSpot *ParkingSpot) (err error){
	fmt.Println("inside repo/AddParkingSpot ******* ", parkingSpot)
	resp := repo.db.Create(&parkingSpot)
//	return nil
	if resp.Error != nil{
		return resp.Error
	}
	return nil
}

func (repo *OwnerRepo)DeleteParkingSpot(ctx context.Context, spotID uint) (err error){
	resp := repo.db.Delete(&ParkingSpot{}, spotID)
	if resp.Error != nil{
		return resp.Error
	}
	return nil
}

func (repo *OwnerRepo)UpdateParkingSpot(ctx context.Context, parkingSpotData ParkingSpot)(parkingSpot ParkingSpot, err error){
	fmt.Println("inside UpdateParkingSpot ***** ", parkingSpotData, parkingSpotData.SpotID)
	if err := repo.db.Model(&parkingSpot).Where("spot_id = ?", parkingSpotData.SpotID).Update(&parkingSpotData).Error; err != nil {
		fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%")
		if gorm.IsRecordNotFoundError(err){
			fmt.Println("RecordNotFound")
			repo.db.Create(&parkingSpotData)  // create new record from newUser
		}
	}
	fmt.Println("Print parkingspot updated ***********", parkingSpot )
	return parkingSpot, err
}
