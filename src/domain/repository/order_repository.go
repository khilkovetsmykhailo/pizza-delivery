package domain

import (
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderRepository interface {
	Save(order domain.Order) (domain.Order, error)
	CheckOrderStatus(orderId primitive.ObjectID) (domain.OrderStatus, error)
	CancelOrder(orderId primitive.ObjectID) (*domain.Order, error)
	GetById(orderId primitive.ObjectID) (*domain.Order, error)
	Update(order domain.Order) (*domain.Order, error)
}
