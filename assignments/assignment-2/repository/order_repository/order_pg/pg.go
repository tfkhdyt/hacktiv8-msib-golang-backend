package order_pg

import (
	"assignment_2/entity"
	"assignment_2/pkg/errs"
	"assignment_2/repository/order_repository"

	"gorm.io/gorm"
)

type orderPG struct {
	db *gorm.DB
}

func NewOrderPG(db *gorm.DB) order_repository.OrderRepository {
	return &orderPG{
		db: db,
	}
}

func (o *orderPG) CreateOrder(
	orderPayload entity.Order,
	itemsPayload []entity.Item,
) (*entity.Order, errs.MessageErr) {
	if err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&orderPayload).Error; err != nil {
			return err
		}

		for _, item := range itemsPayload {
			if err := tx.Model(&orderPayload).Association("Items").Append(&item); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, errs.NewBadRequest(err.Error())
	}

	return &orderPayload, nil
}

func (o *orderPG) GetAllOrders() ([]entity.Order, errs.MessageErr) {
	var orders []entity.Order

	if err := o.db.Preload("Items").Find(&orders).Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to get all orders")
	}

	return orders, nil
}
