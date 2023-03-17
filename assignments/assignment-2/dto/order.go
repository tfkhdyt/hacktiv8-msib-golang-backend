package dto

import "time"

type NewOrderRequest struct {
	OrderedAt    time.Time        `json:"orderedAt"       binding:"required"`
	CustomerName string           `json:"customerName"    binding:"required"`
	Items        []NewItemRequest `json:"items,omitempty" binding:"required"`
}

type NewOrderResponse struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       NewOrderRequest `json:"data"`
}

type GetAllOrdersResponse struct {
	Id           uint                  `json:"id"`
	CreatedAt    time.Time             `json:"createdAt"`
	UpdatedAt    time.Time             `json:"updatedAt"`
	CustomerName string                `json:"customerName"`
	Items        []GetAllItemsResponse `json:"items"`
}
