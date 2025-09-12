package handeler

import (
	pb "account/account_proto"
	// model "account/db"
	"account/service"
	"context"
)

type AccountHandeler struct{
	pb.UnimplementedAccountServiceServer
	service service.AccountService
}

func NewHandeler(service service.AccountService ) *AccountHandeler{
	return  &AccountHandeler{service: service}
}


func (h *AccountHandeler) CreateAccount(ctx context.Context, req *pb.AccountRequest) (*pb.AccountResponse, error) {
	userName:=req.UserName
	password:=req.Password
	amt:=req.Amt
	newAcc,err:=h.service.CreateAccount(userName,amt,password)
	return &pb.AccountResponse{AccId: newAcc.Acc_id},err

}
func (h *AccountHandeler) ReadAccount(ctx context.Context, req *pb.ReadAccountRequest) (*pb.ReadAccountResponse, error){
	id:=req.AccId
	newAcc,err:=h.service.ReadAccount(id)
	return &pb.ReadAccountResponse{AccId: newAcc.Acc_id,Amt: newAcc.Amt,UserName: newAcc.UserName},err
}
func (h *AccountHandeler) UpdateAccount (ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error){
	UserName:=req.UserName
	amt:=req.Amt
	id:=req.AccId
	newAcc,err:=h.service.UpdateAccount(UserName,amt,id)
	return &pb.UpdateAccountResponse{AccId: newAcc.Acc_id,UserName: newAcc.UserName,Amt: newAcc.Amt},err
}
func (h *AccountHandeler) DeleteAccount (ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error){
	id:=req.AccId
 err:=h.service.DeleteAccount(id)
	return &pb.DeleteAccountResponse{Msg: "Deleted"},err

}