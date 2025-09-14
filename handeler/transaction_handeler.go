package handeler

import (
	"account/service"
	pb "account/transaction_proto"
	"context"
)


type TransactionHandeler struct {
	pb.UnimplementedTransactionServiceServer
	service service.TransactionService
}



func NewTransactionHandeler(service service.TransactionService) *TransactionHandeler {
	return &TransactionHandeler{service: service}
}

func (s *TransactionHandeler) CreateTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	from := req.From
	to := req.To
	remark := req.Remark
	amt := req.Amt
	newTrans, err := s.service.Create(from, to, remark, amt)
	return &pb.TransactionResponse{TransId: newTrans.Trans_id}, err
}

func (s *TransactionHandeler) ReadTransaction(ctx context.Context, req *pb.ReadTransactionRequest) (*pb.ReadTransactionResponse, error) {
	id := req.TransId
	newTrans, err := s.service.Read(id)
	return &pb.ReadTransactionResponse{From: newTrans.From, To: newTrans.To, Remark: newTrans.Remark, Amt: newTrans.Amt}, err
}
