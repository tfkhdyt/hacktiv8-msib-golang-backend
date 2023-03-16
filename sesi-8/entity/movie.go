package entity

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title string
	Price uint
}
