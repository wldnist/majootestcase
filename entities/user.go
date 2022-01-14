package entities

type User struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"-"`
	Name     string `gorm:"type:varchar(45)" json:"-"`
	UserName string `gorm:"type:varchar(45);unique;" json:"-"`
	Password string `gorm:"type:varchar(225)" json:"-"`
}
