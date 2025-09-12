package service

import (
	model "account/db"
	"account/repo"
)

type AccountService interface {
	CreateAccount(userName string,amt int64,password string) (*model.Account,error)
	ReadAccount(id int64) (*model.Account,error)
	UpdateAccount(userName string , amt int64,id int64 ) (*model.Account,error)
	DeleteAccount(id int64 ) error
}
type accountService struct{
	repo repo.AccountRepo
}

func NewService(repo repo.AccountRepo) AccountService{
	return  &accountService{repo:repo}
}
func (s *accountService) CreateAccount(userName string,amt int64,password string) (*model.Account,error){
	newAccount:=&model.Account{
		UserName: userName,
		Amt: amt,
		Password: password,
	}
	err:=s.repo.Create(newAccount)
	return newAccount,err
}

func (s *accountService)  ReadAccount(id int64) (*model.Account,error) {
	newAccount:=&model.Account{}
	err:=s.repo.Read(id,newAccount)
	return  newAccount,err
}

func (s *accountService) UpdateAccount(userName string , amt int64 ,id int64) (*model.Account,error)  {
	newAccount:=&model.Account{Acc_id: id}
	err:=s.repo.Update(userName,amt,newAccount)
	return  newAccount,err
}
func (s*accountService) DeleteAccount(id int64 ) error{
   return s.repo.Delete(id) 
}