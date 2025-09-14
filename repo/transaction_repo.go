package repo

import (
	model "account/db"

	"gorm.io/gorm"
)

type TransactionRepo interface{
	Create(newAcc *model.Transaction ) error
	Read(newAcc *model.Transaction ,id int64) error
}
type transactionRepo struct{
		db *gorm.DB
}
func NewTransactionRepo(db *gorm.DB) (transactionRepo){
	return  transactionRepo{db:db}
}

func (t *transactionRepo) Create(newTrans *model.Transaction ) error  {
 err:=t.db.Create(&newTrans).Error
 return  err
}
func (t *transactionRepo) Read(newTrans *model.Transaction ,id int64) error{
	err:=t.db.First(&newTrans,id).Error
	return err


}