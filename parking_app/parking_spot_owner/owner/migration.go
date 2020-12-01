package owner

import(
	"fmt"
	"github.com/jinzhu/gorm"
)

type migrationStruct struct{
	db *gorm.DB
}

func NewMigration(db *gorm.DB){
	fmt.Println("Inside NewMigration *** ")
	m := &migrationStruct{
		db: db,
	}
	m.Migrate()
}

func (m *migrationStruct)Migrate(){
	m.db.AutoMigrate(&OwnerDetails{}, &ParkingSpot{}, &VehicleDetails{})
}