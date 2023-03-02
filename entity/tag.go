package entity

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Tipe string `gorm:"type:VARCHAR(20);NOT NULL"`
}
