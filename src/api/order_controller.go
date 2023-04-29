package api

import (
	"github.com/gin-gonic/gin"
	"github.com/marijakljestan/golang-web-app/src/api/dto"
	"github.com/marijakljestan/golang-web-app/src/domain/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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

	orderDto.CustomerUsername = ctx.GetString("username")
	order, err := handler.orderService.CreateOrder(orderDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"order": order})
}

func (handler *OrderController) CheckOrderStatus(ctx *gin.Context) {
	orderId, convErr := primitive.ObjectIDFromHex(ctx.Param("id"))
	if convErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id provided!"})
		return
	}

	orderStatus, err := handler.orderService.CheckOrderStatus(orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id provided!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": orderStatus.String()})
}

func (handler *OrderController) CancelOrder(ctx *gin.Context) {
	orderId, convErr := primitive.ObjectIDFromHex(ctx.Param("id"))
	if convErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id provided!"})
		return
	}

	order, err := handler.orderService.CancelOrder(orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Order can't  be cancelled"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"order": order})
}

func (handler *OrderController) CancelOrderRegardlessStatus(ctx *gin.Context) {
	orderId, convErr := primitive.ObjectIDFromHex(ctx.Param("id"))
	if convErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id provided!"})
		return
	}

	order, err := handler.orderService.CancelOrderRegardlessStatus(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"order": order})
}
