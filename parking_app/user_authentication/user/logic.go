package user

import(
	"context"
	"errors"
	"fmt"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("gosecretkey")

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

func getHash(pwd []byte) string {
	fmt.Println("Inside getHash *********** ")
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}

func GenerateJWT()(string, error){
	log.Println("Inside GenerateJWT ******* ")
	token:= jwt.New(jwt.SigningMethodHS256)
	tokenString, err :=  token.SignedString(SECRET_KEY)
	if err !=nil{
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
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

func (serv *service)Register(ctx context.Context, user *User) (usr *User, err error){
	fmt.Println("Inside logic/Register *********** ", user)
	user.Password = getHash([]byte(user.Password))
	fmt.Println("after getHash *********** ", user)
	usr, err = serv.repo.Register(ctx, user)
	if err != nil{
		return nil, err
	}
	fmt.Println("repo response *********** ", usr)
	return usr, nil
}


func (serv *service)Login(ctx context.Context, user *User) (*LoginDetails, error){
	fmt.Println("Inside Login **************** ")
	resp, err := serv.repo.Login(ctx, user)
	//var loginDetails *LoginDetails
	fmt.Println("Response from DB ***** ", resp)
	if err != nil{
		log.Println("Database error : ", err)
		return nil, err
	}

	userPass:= []byte(user.Password)
	dbPass:= []byte(resp.Password)
	  
	passErr:= bcrypt.CompareHashAndPassword(dbPass, userPass)

  	if passErr != nil{
	  log.Println("wrong Password : ", passErr)
	  return nil, passErr
  	}
	jwtToken, err := GenerateJWT()
	
	if err != nil{
		return nil, err
	}
	loginDetails := &LoginDetails{
		Token: jwtToken,
	} 
	loginDetails.Token = jwtToken
	fmt.Println("Print loginDetails ***** ", loginDetails)
	return loginDetails, nil
	
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
