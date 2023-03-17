package service

import (
	"assignment_2/dto"
	"assignment_2/entity"
	"assignment_2/pkg/errs"
	"assignment_2/repository/order_repository"
	"net/http"
)

type OrderService interface {
	CreateOrder(payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr)
}

type orderService struct {
	orderRepo order_repository.OrderRepository
}

func NewOrderService(orderRepo order_repository.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) CreateOrder(
	payload dto.NewOrderRequest,
) (*dto.NewOrderResponse, errs.MessageErr) {
	orderPayload := entity.Order{
		CustomerName: payload.CustomerName,
		OrderedAt:    payload.OrderedAt,
	}

	itemsPayload := []entity.Item{}
	for _, eachItem := range payload.Items {
		item := entity.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
		}

		itemsPayload = append(itemsPayload, item)
	}

	newOrder, err := o.orderRepo.CreateOrder(orderPayload, itemsPayload)
	if err != nil {
		return nil, err
	}

	response := &dto.NewOrderResponse{
		Message:    "success",
		StatusCode: http.StatusCreated,
		Data: dto.NewOrderRequest{
			OrderedAt:    newOrder.OrderedAt,
			CustomerName: newOrder.CustomerName,
		},
	}

	return response, nil
}
