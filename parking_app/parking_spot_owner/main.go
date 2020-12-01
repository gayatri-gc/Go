package main

import(
	"context"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"parking_slot_owner/owner"
)

const addr = "postgresql://park@localhost:26257/parkinfo?sslmode=disable"

func main(){
	fmt.Println("Inside main ** ")
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.LogMode(true)
	ctx := context.Background()
	fmt.Println("Start DI ******* ")
	owner.NewMigration(db)
	//owner.NewMigration(db)
	//owner.Migrate()
	repo := owner.NewRepository(db)
	fmt.Println("after call of NewRepository ********")
	serv := owner.NewService(repo)
	endpoints := owner.MakeEndpoints(serv)
	
	fmt.Println("Get endpoints ********** ", endpoints)

	errs := make(chan error)

	go func(){
		handler := owner.NewHttpServer(ctx, endpoints)
		errs <- http.ListenAndServe(":8083", handler)
	}()

	fmt.Println("Error : ", <- errs)
}
