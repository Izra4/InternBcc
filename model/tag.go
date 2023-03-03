package model

import (
	"InternBCC/database"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Tipe string `json:"tipe" binding:"required"`
}

func TagDummy() {
	tag1 := Tag{Tipe: "nikahan"}
	tag2 := Tag{Tipe: "konser"}
	tag3 := Tag{Tipe: "seminar"}
	if err := database.DB.Create(&tag1).Error; err != nil {
		return
	}
	if err := database.DB.Create(&tag2).Error; err != nil {
		return
	}
	if err := database.DB.Create(&tag3).Error; err != nil {
		return
	}
}
