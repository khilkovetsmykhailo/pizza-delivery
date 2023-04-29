package in_memory_repository

import (
	"errors"
	domain "github.com/marijakljestan/golang-web-app/server/domain/model"
	repository "github.com/marijakljestan/golang-web-app/server/domain/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderInmemoryRepository struct{}

var orders []domain.Order

func NewOrderInmemoryRepository() repository.OrderRepository {
	return &OrderInmemoryRepository{}
}

func (repository *OrderInmemoryRepository) Save(order domain.Order) (domain.Order, error) {
	order.Id = primitive.NewObjectID()
	orders = append(orders, order)
	return order, nil
}

func (repository *OrderInmemoryRepository) CheckOrderStatus(orderId primitive.ObjectID) (domain.OrderStatus, error) {
	for _, v := range orders {
		if v.Id == orderId {
			return v.Status, nil
		}
	}
	return -1, errors.New("order with provided id does not exist")
}

func (repository *OrderInmemoryRepository) CancelOrder(orderId primitive.ObjectID) (*domain.Order, error) {
	var order domain.Order
	for i, v := range orders {
		if v.Id == orderId {
			(&orders[i]).Status = domain.CANCELLED
			order = orders[i]
			return &order, nil
		}
	}
	return &order, errors.New("order with provided id does not exist")
}

func (repository *OrderInmemoryRepository) GetById(orderId primitive.ObjectID) (*domain.Order, error) {
	var order domain.Order
	for _, v := range orders {
		if v.Id == orderId {
			order = v
			return &order, nil
		}
	}
	return &order, errors.New("order with provided id does not exist")
}

func (repository *OrderInmemoryRepository) Update(order domain.Order) (*domain.Order, error) {
	for i, v := range orders {
		if v.Id == order.Id {
			orders[i] = order
			return &order, nil
		}
	}
	return &order, errors.New("order with provided id does not exist")
}

func (repository *OrderInmemoryRepository) DeleteAll() {
	orders = []domain.Order{}
}
