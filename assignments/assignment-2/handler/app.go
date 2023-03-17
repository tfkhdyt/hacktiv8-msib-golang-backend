package handler

import (
	"log"

	"assignment_2/database"
	"assignment_2/handler/http_handler"
	"assignment_2/repository/order_repository/order_pg"
	"assignment_2/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	db := database.GetPostgresInstance()

	orderRepo := order_pg.NewOrderPG(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := http_handler.NewOrderHandler(orderService)

	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)

	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err.Error())
	}
}
