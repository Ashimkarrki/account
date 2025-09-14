package main

import (
	pb_acc "account/account_proto"
	pb_tans "account/transaction_proto"
	"os"

	model "account/db"
	"account/handeler"
	"account/repo"
	"account/service"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func main() {
	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port:=os.Getenv("GRPC_PORT")
	

	db, err := gorm.Open(sqlite.Open("xyz.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection to the db failed")
	}
	db.AutoMigrate(&model.Account{},&model.Transaction{})

	lis,err:=net.Listen("tcp",":"+port)
	if err!=nil{
		log.Fatalf("Error in connection to server %v",err)
	}
	accRepo := repo.NewRepo(db)
	transRepo:=repo.NewTransactionRepo(db)

    accService := service.NewService(accRepo)
	transService :=service.NewTransactionService(&transRepo)


    accHandler := handeler.NewHandeler(accService)
	transHandler :=handeler.NewTransactionHandeler(transService)


	grpcServer := grpc.NewServer()
	pb_acc.RegisterAccountServiceServer(grpcServer,accHandler)
	pb_tans.RegisterTransactionServiceServer(grpcServer,transHandler)
	fmt.Printf("Server running on port :%s",port)
	if err:=grpcServer.Serve(lis); err!=nil{
				log.Fatalf("failed to serve: %v", err)
	}

}