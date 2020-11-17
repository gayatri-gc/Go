package main

import(
	"context"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"parking_app/user"
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
	repo := user.NewRepository(db)
	fmt.Println("after call of NewRepository ********")
	serv := user.NewService(repo)
	endpoints := user.MakeEndpoints(serv)
	
	errs := make(chan error)

	go func(){
		handler := user.NewHttpServer(ctx, endpoints)
		errs <- http.ListenAndServe(":8081", handler)
	}()

	fmt.Println("Error : ", <- errs)
}
