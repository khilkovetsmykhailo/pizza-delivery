package domain

import domain "github.com/marijakljestan/golang-web-app/src/domain/model"

type OrderRepository interface {
	CreateOrder(order domain.Order) (domain.Order, error)
	CheckOrderStatus(orderId int) (domain.OrderStatus, error)
	CancelOrder(orderId int) (domain.Order, error)
	GetById(orderId int) (domain.Order, error)
}
