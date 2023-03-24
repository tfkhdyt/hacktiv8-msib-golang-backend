package order_pg

import (
	"assignment_2/entity"
	"assignment_2/pkg/errs"
	"assignment_2/repository/item_repository"
	"assignment_2/repository/order_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type orderPG struct {
	db       *gorm.DB
	itemRepo item_repository.ItemRepository
}

func NewOrderPG(db *gorm.DB, itemRepo item_repository.ItemRepository) order_repository.OrderRepository {
	return &orderPG{
		db:       db,
		itemRepo: itemRepo,
	}
}

func (o *orderPG) CreateOrder(orderPayload *entity.Order, itemsPayload []entity.Item) (*entity.Order, errs.MessageErr) {
	if err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(orderPayload).Error; err != nil {
			return err
		}

		for _, item := range itemsPayload {
			if err := tx.Model(orderPayload).Association("Items").Append(item); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, errs.NewBadRequest(err.Error())
	}

	return orderPayload, nil
}

func (o *orderPG) GetAllOrders() ([]entity.Order, errs.MessageErr) {
	var orders []entity.Order

	if err := o.db.Preload("Items").Find(&orders).Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to get all orders")
	}

	return orders, nil
}

func (o *orderPG) GetOrderByID(orderID uint) (*entity.Order, errs.MessageErr) {
	var order entity.Order

	if err := o.db.Preload("Items").First(&order, orderID).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Order with id %d is not found", orderID))
	}

	return &order, nil
}

func (o *orderPG) UpdateOrderByID(orderID uint, orderPayload *entity.Order, itemsPayload []entity.Item) (*entity.Order, errs.MessageErr) {
	order, err := o.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	tx := o.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Panicf("Error: %v\n", err.Error())
		return nil, errs.NewInternalServerError("Failed to begin transaction")
	}

	if err := tx.Model(order).Updates(orderPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewBadRequest(fmt.Sprintf("Order with id %d failed to update. %v", orderID, err.Error()))
	}

	order.Items = []entity.Item{}
	for _, item := range itemsPayload {
		updatedItem, err := o.itemRepo.UpdateItemByItemCode(item.ItemCode, &item, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		order.Items = append(order.Items, *updatedItem)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to commit transaction")
	}

	return order, nil
}

func (o *orderPG) DeleteOrderByID(orderID uint) errs.MessageErr {
	order, err := o.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	if err := o.db.Delete(order).Error; err != nil {
		return errs.NewInternalServerError("Failed to delete order")
	}

	return nil
}
