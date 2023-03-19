package handler

import (
	"assignment_2/database"
	"assignment_2/docs"
	"assignment_2/handler/http_handler"
	"assignment_2/repository/order_repository/order_pg"
	"assignment_2/service"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {
	docs.SwaggerInfo.Title = "Order and Item API"
	docs.SwaggerInfo.Description = "This is an assignment for Hacktiv8 Golang Course"
	docs.SwaggerInfo.Version = "0.3"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := database.GetPostgresInstance()

	orderRepo := order_pg.NewOrderPG(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := http_handler.NewOrderHandler(orderService)

	r := gin.Default()

	// Create
	r.POST("/orders", orderHandler.CreateOrder)
	// Get All
	r.GET("/orders", orderHandler.GetAllOrders)
	// Get One
	r.GET("/orders/:orderID", orderHandler.GetOrderByID)
	// Update
	r.PATCH("/orders/:orderID", orderHandler.UpdateOrderByID)
	// Delete
	r.DELETE("/orders/:orderID", orderHandler.DeleteOrderByID)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err.Error())
	}
}
