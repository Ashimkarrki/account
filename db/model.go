package model
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


