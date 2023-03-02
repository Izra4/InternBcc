package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama      string `gorm:"NOT NULL"`
	Email     string `gorm:"unique;NOT NULL"`
	Password  string `gorm:"NOT NULL"`
	Number    string `gorm:"type:varchar(20);NOT NULL"`
	Testimoni Testimoni
}
