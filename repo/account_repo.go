package repo

import (
	model "account/db"

	"gorm.io/gorm"
)
type RepoDb struct{
		db *gorm.DB
}
type AccountRepo interface{
	Create(acc *model.Account) error
	Read(id int64, acc *model.Account)  error
	Update(userName string,amt int64,acc *model.Account) error
	Delete(id int64 ) error
}

func NewRepo(db *gorm.DB) (*RepoDb){
	return  &RepoDb{db:db}
}
func ( r *RepoDb) Create(acc *model.Account) error{
	return r.db.Create(acc).Error
}

func (r *RepoDb) Read(id int64, acc *model.Account) error{
	return r.db.First(acc,id).Error
}
func (r *RepoDb) Update(userName string,amt int64,acc *model.Account) error{
	return r.db.Model(&acc).Updates(model.Account{UserName: userName,Amt: amt}).Error
}
func (r *RepoDb) Delete(id int64 ) error  {
	return r.db.Delete(&model.Account{},id).Error
}