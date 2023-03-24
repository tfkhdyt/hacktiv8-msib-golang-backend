package handler

import (
	"assignment_2/database"
	"assignment_2/docs"
	"assignment_2/handler/http_handler"
	"assignment_2/repository/item_repository/item_pg"
	"assignment_2/repository/order_repository/order_pg"
	"assignment_2/service"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	docs.SwaggerInfo.Title = "Order and Item API"
	docs.SwaggerInfo.Description = "This is an assignment for Hacktiv8 Golang Course"
	docs.SwaggerInfo.Version = "v0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func StartApp() {
	db := database.GetPostgresInstance()

	itemRepo := item_pg.NewItemPG(db)
	orderRepo := order_pg.NewOrderPG(db, itemRepo)
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
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err.Error())
	}
}
