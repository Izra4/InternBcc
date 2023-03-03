package model

import (
	"InternBCC/database"
	"gorm.io/gorm"
)

type Gedung struct {
	gorm.Model
	Nama      string `json:"nama" binding:"required"`
	Fasilitas string `json:"fasilitas" binding:"required"`
}

func GDummy() {
	G1 := Gedung{
		Model:     gorm.Model{},
		Nama:      "Kartika Graha",
		Fasilitas: "100 Kursi, Dobly sound system",
	}
	G2 := Gedung{
		Model:     gorm.Model{},
		Nama:      "Singha Graha",
		Fasilitas: "150 Kursi, Dobly sound system",
	}
	if err := database.DB.Create(&G1).Error; err != nil {
		return
	}
	if err := database.DB.Create(&G2).Error; err != nil {
		return
	}
}
