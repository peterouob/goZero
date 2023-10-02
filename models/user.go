package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity string `gorm:"column:identity;type varchar(50)" json:"identity"`
	Name     string `gorm:"column:name;type varchar(50)" json:"name"`
	Password string `gorm:"column:password;type varchar(50)" json:"password"`
}

func (u *User) TableName() string {
	return "user"
}
