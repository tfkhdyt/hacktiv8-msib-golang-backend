package dto

import (
	"assignment_2/entity"
	"time"
)

type NewOrderRequest struct {
	OrderedAt    time.Time        `json:"orderedAt" example:"2023-03-19T18:55:00+07:00"`
	CustomerName string           `json:"customerName" binding:"required" example:"Taufik Hidayat"`
	Items        []NewItemRequest `json:"items" binding:"required"`
}

func (o *NewOrderRequest) OrderRequestToEntity() *entity.Order {
	return &entity.Order{
		CustomerName: o.CustomerName,
		OrderedAt:    o.OrderedAt,
	}
}

type NewOrderResponse struct {
	StatusCode int             `json:"statusCode" example:"201"`
	Message    string          `json:"message" example:"Order with id 69 has been successfully created"`
	Data       NewOrderRequest `json:"data"`
}

type GetAllOrdersResponse struct {
	StatusCode int         `json:"statusCode" example:"200"`
	Message    string      `json:"message" example:"success"`
	Data       []OrderData `json:"data"`
}

type GetOrderByIDResponse struct {
	StatusCode int       `json:"statusCode" example:"200"`
	Message    string    `json:"message" example:"success"`
	Data       OrderData `json:"data"`
}

type DeleteOrderByIDResponse struct {
	StatusCode int    `json:"statusCode" example:"200"`
	Message    string `json:"message" example:"Order with id 69 has been successfully deleted"`
}

type OrderData struct {
	ID           uint       `json:"id" example:"1"`
	CreatedAt    time.Time  `json:"createdAt" example:"2023-03-19T18:55:00+07:00"`
	UpdatedAt    time.Time  `json:"updatedAt" example:"2023-03-19T18:55:00+07:00"`
	CustomerName string     `json:"customerName" example:"Taufik Hidayat"`
	Items        []ItemData `json:"items"`
}
