package item_pg

import (
	"assignment_2/entity"
	"assignment_2/pkg/errs"
	"assignment_2/repository/item_repository"
	"fmt"

	"gorm.io/gorm"
)

type itemPG struct {
	db *gorm.DB
}

func NewItemPG(db *gorm.DB) item_repository.ItemRepository {
	return &itemPG{
		db: db,
	}
}

func (i *itemPG) GetItemByItemCode(itemCode string, txs ...*gorm.DB) (*entity.Item, errs.MessageErr) {
	var item entity.Item

	if len(txs) > 0 {
		tx := txs[0]

		if err := tx.First(&item, "item_code = ?", itemCode).Error; err != nil {
			return nil, errs.NewNotFound(fmt.Sprintf("Item with item code %s is not found", itemCode))
		}

		return &item, nil
	}

	if err := i.db.First(&item, "item_code = ?", itemCode).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Item with item code %s is not found", itemCode))
	}

	return &item, nil
}

func (i *itemPG) UpdateItemByItemCode(itemCode string, payload *entity.Item, txs ...*gorm.DB) (*entity.Item, errs.MessageErr) {
	item, err := i.GetItemByItemCode(itemCode, txs...)
	if err != nil {
		return nil, err
	}

	if len(txs) > 0 {
		tx := txs[0]
		if err := tx.Model(item).Updates(payload).Error; err != nil {
			return nil, errs.NewBadRequest(fmt.Sprintf("Item with item code %s failed to update. %v", itemCode, err.Error()))
		}

		return item, nil
	}

	if err := i.db.Model(item).Updates(payload).Error; err != nil {
		return nil, errs.NewBadRequest(fmt.Sprintf("Item with item code %s failed to update. %v", itemCode, err.Error()))
	}

	return item, nil
}
