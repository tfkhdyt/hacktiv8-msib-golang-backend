package order_pg

import (
	"assignment_2/entity"
	"assignment_2/pkg/errs"
	"assignment_2/repository/order_repository"
	"log"

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
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed when creating order!")
	}

	return &orderPayload, nil
}
