package main

import (
	"fmt"
	"log"
	"os"

	pb_acc "account/account_proto"
	pb_trans "account/transaction_proto"

	"account/rest"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)




func main(){
		err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port:=os.Getenv("REST_PORT")
	con,err:=grpc.NewClient("localhost:"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
		fmt.Printf("Failed to connect: %v \n", err)
	}
	defer con.Close()
	clientAcc:=pb_acc.NewAccountServiceClient(con)
	clientTrans:=pb_trans.NewTransactionServiceClient(con)


	accHandler:=rest.AccountHandelers{Client:clientAcc}
	transHandler:=rest.Transactionhandeler{Client:clientTrans}

	e:=echo.New()
	e.GET("/account/:id",accHandler.GetAccount)
	e.POST("/account",accHandler.AdddAccount)
	e.PUT("/account/:id",accHandler.UpdateAccount)
	e.DELETE("/account/:id",accHandler.DeleteAccount)
	e.POST("/transaction",transHandler.CreateTransaction)
	e.GET("/transaction/:id",transHandler.ReadTransaction)



fmt.Printf("Running on port :%s",port)
	e.Start(":"+port)
	

}