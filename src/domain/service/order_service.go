package service

import (
	"fmt"
	model "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
)

type OrderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(pizzaRepository repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: pizzaRepository,
	}
}

func (service *OrderService) ListMenu() ([]model.Pizza, error) {
	menu, err := service.orderRepository.GetMenu()
	if err != nil {
		fmt.Println(err)
	}

	return menu, nil
}
