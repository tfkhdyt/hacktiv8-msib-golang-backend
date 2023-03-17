package dto

import "time"

type NewOrderRequest struct {
	OrderedAt    time.Time        `json:"orderedAt"`
	CustomerName string           `json:"customerName"`
	Items        []NewItemRequest `json:"items"`
}

type NewOrderResponse struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       NewOrderRequest `json:"data"`
}
