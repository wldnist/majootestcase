package entities

type Merchant struct {
	ID           int64  `gorm:"primary_key:auto_increment" json:"-"`
	UserID       int64  `gorm:"type:bigint;not null;" json:"-"`
	MerchantName string `gorm:"type:varchar(45);not null;" json:"-"`
}
