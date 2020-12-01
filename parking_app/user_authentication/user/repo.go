package user

import 
(
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
)

const ExpectedMinTime = 120

type UserRepo struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository{
	fmt.Println("Inside NewRepository ********", db)
	return &UserRepo{
		db: db,
	}
} 

func (repo UserRepo)Register(ctx context.Context, user *User) (usr *User, err error){
	fmt.Println("Inside repo/Register *********** ")
	//check if user already exists, else create
	//  d := repo.db.Find(&user, "email = ?", user.Email)
	//  fmt.Println("check : ", d)
	// userdata := repo.db.Where("email=?", user.Email).Find(&usr)
	// fmt.Println("Print userData **** ", userdata)
	// row := repo.db.Table("users").Where("email = ?", user.Email).Row()
	// row.Scan(&user.Email)
	// fmt.Println("Scan rows ******* ", row)
	result := repo.db.Create(&user)
	if result.Error != nil{
		return nil, result.Error
	}
	if result.RowsAffected > 0{
		usr = user
	}
	return usr, nil
}

func (repo UserRepo)Login(ctx context.Context, user *User) (usr User, err error){
	userdata := repo.db.Where("email=?", user.Email).First(&usr)
	if userdata.Error != nil {
		fmt.Println("Inside Login when gets userdata**************** ", userdata.Error)
		return usr,  userdata.Error
	}
	return usr, nil
}

func(repo UserRepo)ParkVehicle(ctx context.Context, Vdata VehicleDetails) error{
	fmt.Println("Inside CreateUser ********* ", Vdata)
	repo.db.AutoMigrate(&VehicleDetails{})
	_ = repo.db.Create(&Vdata)
	return nil

}