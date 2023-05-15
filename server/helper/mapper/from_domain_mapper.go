package mapper

import (
	"github.com/marijakljestan/golang-web-app/server/api/dto"
	domain "github.com/marijakljestan/golang-web-app/server/domain/model"
)

func MapOrderFromDomain(order domain.Order) dto.OrderDto {
	orderDto := dto.OrderDto{
		Id:               order.Id.Hex(),
		CustomerUsername: order.CustomerUsername,
		Price:            order.Price,
		Status:           order.Status.String(),
		Items:            []dto.OrderItemDto{},
	}
	for _, item := range order.Items {
		orderDto.Items = append(orderDto.Items, MapOrderItemFromDomain(item))
	}
	return orderDto
}

func MapOrderItemFromDomain(orderItem domain.OrderItem) dto.OrderItemDto {
	orderItemDto := dto.OrderItemDto{
		PizzaName: orderItem.PizzaName,
		Quantity:  orderItem.Quantity,
	}
	return orderItemDto
}
