package order_repository

import (
	"assignment_2/entity"
	"assignment_2/pkg/errs"
)

type OrderRepository interface {
	CreateOrder(
		orderPayload entity.Order,
		itemPayloads []entity.Item,
	) (*entity.Order, errs.MessageErr)
}
