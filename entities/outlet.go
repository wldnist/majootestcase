package entities

type Outlet struct {
	ID         int64    `gorm:"primary_key:auto_increment" json:"-"`
	MerchantID int64    `gorm:"type:bigint;not null;" json:"-"`
	OutletName string   `gorm:"type:varchar(45);not null;" json:"-"`
	Merchant   Merchant `gorm:"foreignkey:MerchantID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
