package persistence

import (
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
)

type OrderInmemoryRepository struct{}

var orders = []domain.Order{}

var orderIdIncrementer int = 0

func NewOrderInmemoryRepository() repository.OrderRepository {
	return &OrderInmemoryRepository{}
}

func (repository *OrderInmemoryRepository) CreateOrder(order domain.Order) (domain.Order, error) {
	orderIdIncrementer++
	order.Id = orderIdIncrementer
	orders = append(orders, order)
	return order, nil
}

func (repository *OrderInmemoryRepository) CheckOrderStatus(orderId int) (domain.OrderStatus, error) {
	for _, v := range orders {
		if v.Id == orderId {
			return v.Status, nil
		}
	}
	return -1, nil
}

func (repository *OrderInmemoryRepository) CancelOrder(orderId int) (domain.Order, error) {
	var order domain.Order
	for _, v := range orders {
		if v.Id == orderId {
			order := v
			order.Status = domain.CANCELLED
			//save order in slice
			break
		}
	}
	return order, nil
}
