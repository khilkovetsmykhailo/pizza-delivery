package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	service "github.com/marijakljestan/golang-web-app/src/domain/service"
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

func (handler *OrderController) GetMenu(ctx *gin.Context) {
	menu, err := handler.orderService.ListMenu()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error getting menu!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": menu})
}
