package mapper

import (
	"github.com/marijakljestan/golang-web-app/server/api/dto"
	model "github.com/marijakljestan/golang-web-app/server/domain/model"
)

func MapPizzaToDomain(pizzaDto dto.PizzaDto) model.Pizza {
	pizza := model.Pizza{
		Name:        pizzaDto.Name,
		Description: pizzaDto.Description,
		Price:       pizzaDto.Price,
	}
	return pizza
}

func MapOrderToDomain(orderDto dto.OrderDto) model.Order {
	order := model.Order{
		CustomerUsername: orderDto.CustomerUsername,
		Items:            []model.OrderItem{},
	}
	for _, v := range orderDto.Items {
		order.Items = append(order.Items, MapOrderItemToDomain(v))
	}
	return order
}

func MapOrderItemToDomain(orderItemDto dto.OrderItemDto) model.OrderItem {
	orderItem := model.OrderItem{
		PizzaName: orderItemDto.PizzaName,
		Quantity:  orderItemDto.Quantity,
	}
	return orderItem
}

func MapUserToDomain(userDto dto.UserDto) model.User {
	user := model.User{
		Username: userDto.Username,
		Password: userDto.Password,
	}
	return user
}
