package dto

import (
	"assignment_2/entity"
	"time"
)

type NewItemRequest struct {
	ItemCode    string `json:"itemCode"    binding:"required" example:"BRNG-001"`
	Description string `json:"description" binding:"required" example:"Ini adalah sebuah barang yang dipesan"`
	Quantity    uint   `json:"quantity"    binding:"required" example:"1"`
}

func (i *NewItemRequest) ItemRequestToEntity() *entity.Item {
	return &entity.Item{
		ItemCode:    i.ItemCode,
		Description: i.Description,
		Quantity:    i.Quantity,
	}
}

type ItemData struct {
	ID          uint      `json:"id" example:"2"`
	CreatedAt   time.Time `json:"createdAt" example:"2023-03-19T18:55:00+07:00"`
	UpdatedAt   time.Time `json:"updatedAt" example:"2023-03-19T18:55:00+07:00"`
	ItemCode    string    `json:"itemCode" example:"BRNG-001"`
	Description string    `json:"description" example:"Ini adalah sebuah barang yang dipesan"`
	Quantity    uint      `json:"quantity" example:"1"`
	OrderID     uint      `json:"orderId" example:"1"`
}

type UpdateItemResponse struct {
	StatusCode int      `json:"statusCode" example:"200"`
	Message    string   `json:"message" example:"Order with itemCode BRH-001 has been successfully updated"`
	Data       ItemData `json:"data"`
}
