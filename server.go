package main

import (
	pb "account/account_proto"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



type AccountServer struct{
	pb.UnimplementedAccountServiceServer
	db *gorm.DB
}

type Account struct {
	Acc_id int64 `gorm:"primaryKey"`
	UserName string `gorm:"uniqueIndex"`
    Amt int64
	Sender Transaction `gorm:"foreignKey:From"`
	Receiver Transaction `gorm:"foreignKey:To"`
	Password string

}
type Transaction struct{
	Trans_id int64 `gorm:"primaryKey"`
    From int64
    To  int64
    Amt int64
    Remark string;
}
func generateJWT(accID int64)(string,error){
	 jwtkey := []byte("supersecretkey")

	claims:=jwt.MapClaims{
		"Acc_id":accID,
		 "exp":      time.Now().Add(time.Hour * 1).Unix(), 
        "iat":      time.Now().Unix(),
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	signed, err := token.SignedString(jwtkey)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}
	return signed,nil
}


func (s *AccountServer) Login (ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error){
	fmt.Println(req.UserName)
		fmt.Println(req.Password)
	acc:=Account{UserName: req.UserName}
	if err:=s.db.Where(Account{}).First(&acc).Error;err!=nil{
		return nil,err
	}

	

	if req.Password!=acc.Password{
		return &pb.LoginResponse{Token: "",Msg: "Incorrect credentials "},nil
	}
	token,err:=generateJWT(acc.Acc_id)
	if err != nil {
		return nil, err
	}
	return  &pb.LoginResponse{Token:token,Msg: "Success" },nil




}

func (s *AccountServer) CreateAccount (ctx context.Context, req *pb.AccountRequest) (*pb.AccountResponse, error){
	newAccount:=Account{
		UserName: req.UserName,
		Amt: req.Amt,
		Password: req.Password,
	}
	fmt.Println(newAccount)
	if err:=s.db.Create(&newAccount).Error; err!=nil{
		return  nil,err
	}
	return &pb.AccountResponse{AccId:newAccount.Acc_id },nil
}
func (s *AccountServer) ReadAccount (ctx context.Context, req *pb.ReadAccountRequest) (*pb.ReadAccountResponse, error){
	id:=req.AccId
	newAccount:=&Account{}
		if err:=s.db.First(&newAccount,id).Error; err!=nil{
			return nil,err
		}

	return &pb.ReadAccountResponse{AccId: newAccount.Acc_id,UserName: newAccount.UserName,Amt: newAccount.Amt},nil
}
func (s *AccountServer) UpdateAccount (ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error){
	UserName:=req.UserName
	amt:=req.Amt
	id:=req.AccId

	newAccount:=&Account{Acc_id: id}
		if err:=s.db.Model(&newAccount).Updates(Account{UserName: UserName,Amt: amt}).Error; err!=nil{
			return nil,err
		}

	return &pb.UpdateAccountResponse{AccId: newAccount.Acc_id,UserName: newAccount.UserName,Amt: newAccount.Amt},nil
}
func (s *AccountServer) DeleteAccount (ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error){
	id:=req.AccId
	newAccount:=&Account{}
		if err:=s.db.Delete(&newAccount,id).Error; err!=nil{
			return nil,err
		}

	return &pb.DeleteAccountResponse{Msg: "Deleted"},nil
}

func (s *AccountServer) CreateTransaction (ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error){
	from:=req.From
	to:=req.To
	remark:=req.Remark
	amt:=req.Amt

	newAccount:=&Transaction{From:from,To: to,Remark: remark,Amt: amt}
		if err:=s.db.Create(&newAccount).Error; err!=nil{
			return nil,err
		}

	return &pb.TransactionResponse{TransId: newAccount.Trans_id},nil
}

func (s *AccountServer) ReadTransaction (ctx context.Context, req *pb.ReadTransactionRequest) (*pb.ReadTransactionResponse, error){
	id:=req.TransId
	newTransaction:=&Transaction{}
		if err:=s.db.First(&newTransaction,id).Error; err!=nil{
			return nil,err
		}

	return &pb.ReadTransactionResponse{From: newTransaction.From,To: newTransaction.To,Remark: newTransaction.Remark},nil
}

func (s *AccountServer) GetAllAccount (ctx context.Context, req *pb.ReadAllAccountRequest) (*pb.ReadAllAccountResponse, error){
		accounts:= []Account{}
	if err:=s.db.Find(&accounts).Error; err!=nil{
		return  nil,err
	}
	accountsList:=&pb.ReadAllAccountResponse{}
	for _,a:=range accounts{
		accountsList.Accounts=append(accountsList.Accounts,&pb.ReadAccountResponse{
			UserName: a.UserName,
			Amt: a.Amt,
			AccId: a.Acc_id,

		})
	}

	return accountsList,nil
}





func main() {

	db, err := gorm.Open(sqlite.Open("xyz.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection to the db failed")
	}
	db.AutoMigrate(&Account{},&Transaction{})
	lis,err:=net.Listen("tcp",":3000")
	if err!=nil{
		log.Fatalf("Error in connection to server %v",err)
	}
	grpcServer:=grpc.NewServer()
	pb.RegisterAccountServiceServer(grpcServer,&AccountServer{db:db})
	fmt.Println("Server running on port :3000")
	if err:=grpcServer.Serve(lis); err!=nil{
				log.Fatalf("failed to serve: %v", err)
	}

}