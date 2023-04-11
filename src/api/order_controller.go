package api

import (
	"github.com/gin-gonic/gin"
	"github.com/marijakljestan/golang-web-app/src/api/dto"
	"github.com/marijakljestan/golang-web-app/src/domain/service"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(service *service.OrderService) *OrderController {
	return &OrderController{
		orderService: service,
	}
}

func (handler *OrderController) CreateOrder(ctx *gin.Context) {
	var orderDto dto.OrderDto
	if err := ctx.BindJSON(&orderDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid arguments passed!"})
		return
	}

	order, err := handler.orderService.CreateOrder(orderDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating order!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": order})
}

func (handler *OrderController) CheckOrderStatus(ctx *gin.Context) {
	orderId, convErr := strconv.Atoi(ctx.Param("id"))
	if convErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id provided!"})
		return
	}
	orderStatus, err := handler.orderService.CheckOrderStatus(orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id provided!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": orderStatus})
}
