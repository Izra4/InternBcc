package entity

import "gorm.io/gorm"

type Testimoni struct {
	gorm.Model
	UserID   uint
	GedungID uint
}
