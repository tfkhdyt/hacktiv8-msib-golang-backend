package http_handler

import (
	"assignment_2/dto"
	"assignment_2/pkg/errs"
	"assignment_2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *orderHandler {
	return &orderHandler{orderService: orderService}
}

// CreateOrder godoc
//
//	@Summary		Create order
//	@Description	Create an order by json
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order	body		dto.NewOrderRequest	true	"Create order request body"
//	@Success		201		{object}	dto.NewOrderResponse
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/orders [post]
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

// GetAllOrders godoc
//
//	@Summary		List orders
//	@Description	Get all orders
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.GetAllOrdersResponse
//	@Failure		500	{object}	errs.MessageErrData
//	@Router			/orders [get]
func (o *orderHandler) GetAllOrders(ctx *gin.Context) {
	orders, err := o.orderService.GetAllOrders()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(orders.StatusCode, orders)
}

// GetOrderByID godoc
//
//	@Summary		Find order
//	@Description	Get an order by id
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint	true	"Order ID"
//	@Success		200	{object}	dto.GetOrderByIDResponse
//	@Failure		400	{object}	errs.MessageErrData
//	@Failure		404	{object}	errs.MessageErrData
//	@Router			/orders/{id} [get]
func (o *orderHandler) GetOrderByID(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		newError := errs.NewBadRequest("orderID should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	order, errOrder := o.orderService.GetOrderByID(uint(orderIDInt))
	if errOrder != nil {
		ctx.JSON(errOrder.StatusCode(), errOrder)
		return
	}

	ctx.JSON(order.StatusCode, order)
}

// UpdateOrderByID godoc
//
//	@Summary		Update order
//	@Description	Update an order by json
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		uint				true	"Order ID"
//	@Param			order	body		dto.NewOrderRequest	true	"Update order request body"
//	@Success		200		{object}	dto.GetOrderByIDResponse
//	@Failure		400		{object}	errs.MessageErrData
//	@Failure		404		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Router			/orders/{id} [patch]
func (o *orderHandler) UpdateOrderByID(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		newError := errs.NewBadRequest("orderID should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	var requestBody dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	updatedOrder, errOrder := o.orderService.UpdateOrderByID(uint(orderIDInt), requestBody)
	if errOrder != nil {
		ctx.JSON(errOrder.StatusCode(), errOrder)
		return
	}

	ctx.JSON(updatedOrder.StatusCode, updatedOrder)
}

// DeleteOrderByID godoc
//
//	@Summary		Delete order
//	@Description	Delete an order by id
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint	true	"Order ID"
//	@Success		200	{object}	dto.DeleteOrderByIDResponse
//	@Failure		400	{object}	errs.MessageErrData
//	@Failure		404	{object}	errs.MessageErrData
//	@Failure		500	{object}	errs.MessageErrData
//	@Router			/orders/{id} [delete]
func (o *orderHandler) DeleteOrderByID(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		newError := errs.NewBadRequest("orderID should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, errOrder := o.orderService.DeleteOrderByID(uint(orderIDInt))
	if errOrder != nil {
		ctx.JSON(errOrder.StatusCode(), errOrder)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
