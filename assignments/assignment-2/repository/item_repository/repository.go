package item_repository

import (
	"assignment_2/entity"
	"assignment_2/pkg/errs"
)

type ItemRepository interface {
	CreateItem(entity.Item) (*entity.Item, errs.MessageErr)
}
