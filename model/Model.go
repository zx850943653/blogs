package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"varchar(50)" json:"name"`
	Sex  string `gorm:"int(2)" json:"sex"`
}
