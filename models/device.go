package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type varchar(50)" json:"identity"`
	ProductIdentity string `gorm:"productidentity:name;type varchar(50)" json:"productidentity"`
	Name            string `gorm:"column:name;type varchar(50)" json:"name"`
	Key             string `gorm:"column:key;type varchar(50)" json:"key"`
	Secret          string `gorm:"column:secret;type varchar(50)" json:"secret"`
}

func (d *Device) TableName() string {
	return "device"
}
