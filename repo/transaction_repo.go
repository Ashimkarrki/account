package repo

import (
	model "account/db"

	"gorm.io/gorm"
)
type TransactionRepo struct{
		db *gorm.DB
}
func NewTransactionRepo(db *gorm.DB) (TransactionRepo){
	return  TransactionRepo{db:db}
}

func (t *TransactionRepo) Create (newAcc *model.Account ) error  {
 err:=t.db.Create(&newAcc).Error
 return  err
}
func (t *TransactionRepo) Read (newAcc *model.Account ,id int64) error{
	err:=t.db.First(&newAcc,id).Error
	return err


}