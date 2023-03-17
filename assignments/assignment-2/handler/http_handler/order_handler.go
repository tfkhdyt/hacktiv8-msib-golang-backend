package http_handler

import (
	"assignment_2/dto"
	"assignment_2/pkg/errs"
	"assignment_2/service"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *orderHandler {
	return &orderHandler{orderService: orderService}
}

func (o *orderHandler) CreateOrder(ctx *gin.Context) {
	var requestBody dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	newOrder, err := o.orderService.CreateOrder(requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(newOrder.StatusCode, newOrder)
}

func (o *orderHandler) GetAllOrders(ctx *gin.Context) {
	orders, err := o.orderService.GetAllOrders()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(orders.StatusCode, orders)
}
