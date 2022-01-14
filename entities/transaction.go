package entities

type Transaction struct {
	ID        int64   `gorm:"primary_key:auto_increment" json:"-"`
	OutletID  int64   `gorm:"type:bigint;not null;" json:"-"`
	BillTotal float64 `gorm:"type:bigint;not null;" json:"-"`
	Outlet    Outlet  `gorm:"foreignkey:OutletID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
