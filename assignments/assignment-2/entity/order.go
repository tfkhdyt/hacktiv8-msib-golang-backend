package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string    `gorm:"not null"`
	OrderedAt    time.Time `gorm:"default:now()"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
