package service

import (
	model "account/db"
	"account/repo"
)
type TransactionService interface{
	Create( from int64 , to int64 ,remark string ,amt int64) (*model.Transaction,error)
	Read(id int64)  (*model.Transaction,error)

}

type transactionService struct{
	repo repo.TransactionRepo
}


func NewTransactionService(repo repo.TransactionRepo) TransactionService  {
	return &transactionService{repo: repo}
}

func (t *transactionService) Create( from int64 , to int64 ,remark string ,amt int64) (*model.Transaction,error) {
	newTrans:=&model.Transaction{
		From: from,
		To: to,
		Remark: remark,
		Amt: amt,
	}
	err:=t.repo.Create(newTrans)
	return newTrans,err
	
}

func (t *transactionService) Read(id int64)   (*model.Transaction,error){
	newTrans:=&model.Transaction{}
	err:=t.repo.Read(newTrans,id)
	return newTrans,err

}