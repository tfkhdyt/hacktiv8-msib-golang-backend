package entity

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ItemCode    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Quantity    uint   `gorm:"not null"`
	OrderID     uint
}
