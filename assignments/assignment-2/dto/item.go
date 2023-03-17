package dto

import "time"

type NewItemRequest struct {
	ItemCode    string `json:"itemCode"    binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity"    binding:"required"`
}

type ItemDTO struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ItemCode    string    `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    uint      `json:"quantity"`
	OrderID     uint      `json:"orderId"`
}
