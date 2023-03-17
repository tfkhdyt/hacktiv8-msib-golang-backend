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
	GetAllOrders() (*dto.GetAllOrdersResponse, errs.MessageErr)
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

func (o *orderService) GetAllOrders() (*dto.GetAllOrdersResponse, errs.MessageErr) {
	orders, err := o.orderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	data := []dto.OrderData{}

	for _, eachOrder := range orders {
		items := []dto.ItemData{}

		for _, eachItem := range eachOrder.Items {
			item := dto.ItemData{
				ID:          eachItem.ID,
				CreatedAt:   eachItem.CreatedAt,
				UpdatedAt:   eachItem.UpdatedAt,
				ItemCode:    eachItem.ItemCode,
				Description: eachItem.Description,
				Quantity:    eachItem.Quantity,
				OrderID:     eachItem.OrderID,
			}

			items = append(items, item)
		}

		order := dto.OrderData{
			ID:           eachOrder.ID,
			CreatedAt:    eachOrder.CreatedAt,
			UpdatedAt:    eachOrder.UpdatedAt,
			CustomerName: eachOrder.CustomerName,
			Items:        items,
		}

		data = append(data, order)
	}

	response := &dto.GetAllOrdersResponse{
		Message:    "success",
		StatusCode: http.StatusOK,
		Data:       data,
	}

	return response, nil
}
