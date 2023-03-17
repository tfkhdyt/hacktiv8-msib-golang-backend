package handler

import (
	"assignment_2/database"
	"assignment_2/repository/order_repository/order_pg"
)

func StartApp() {
	db := database.GetPostgresInstance()

	orderRepo := order_pg.NewOrderPG(db)
}
