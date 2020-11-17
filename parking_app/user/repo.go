package user

import 
(
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	
)

type UserRepo struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository{
	fmt.Println("Inside NewRepository ********", db)
	return &UserRepo{
		db: db,
	}
} 

func(repo UserRepo)ParkVehicle(ctx context.Context, Vdata VehicleDetails) error{
	fmt.Println("Inside CreateUser ********* ", Vdata)
	repo.db.AutoMigrate(&VehicleDetails{})
	_ = repo.db.Create(&Vdata)
	return nil

}