package in_memory_repository

import (
	"errors"
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
)

type OrderInmemoryRepository struct{}

var orders []domain.Order
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
	return -1, errors.New("order with provided id does not exist")
}

func (repository *OrderInmemoryRepository) CancelOrder(orderId int) (domain.Order, error) {
	var order domain.Order
	for i, v := range orders {
		if v.Id == orderId {
			(&orders[i]).Status = domain.CANCELLED
			order = orders[i]
			return order, nil
		}
	}
	return order, errors.New("order with provided id does not exist")
}

func (repository *OrderInmemoryRepository) GetById(orderId int) (domain.Order, error) {
	var order domain.Order
	for _, v := range orders {
		if v.Id == orderId {
			order = v
			return order, nil
		}
	}
	return order, errors.New("order with provided id does not exist")
}

func (repository *OrderInmemoryRepository) UpdateOrder(order domain.Order) (domain.Order, error) {
	for i, v := range orders {
		if v.Id == order.Id {
			orders[i] = order
			return order, nil
		}
	}
	return order, errors.New("order with provided id does not exist")
}
