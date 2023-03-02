package entity

import "gorm.io/gorm"

type Gedung struct {
	gorm.Model
	Nama      string `gorm:"type:VARCHAR(50)"`
	Fasilitas string
	Tag       []Tag
}
