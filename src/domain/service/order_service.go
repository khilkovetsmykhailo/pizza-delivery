package service

import (
	"fmt"
	"github.com/marijakljestan/golang-web-app/src/api/dto"
	"github.com/marijakljestan/golang-web-app/src/domain/mapper"
	model "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
	"time"
)

type OrderService struct {
	orderRepository repository.OrderRepository
	pizzaService    *PizzaService
}

func NewOrderService(orderRepository repository.OrderRepository, pizzaService *PizzaService) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		pizzaService:    pizzaService,
	}
}

func (service *OrderService) CreateOrder(orderDto dto.OrderDto) (model.Order, error) {
	createdOrder := service.initializeAndSaveOrder(orderDto)

	ch := make(chan model.OrderStatus)
	go func(ch chan<- model.OrderStatus) {
		dur := 15 * time.Second
		time.Sleep(dur)
		if orderStatus, _ := service.CheckOrderStatus(createdOrder.Id); orderStatus != model.CANCELLED {
			ch <- model.READY_TO_BE_DELIVERED
		}
		close(ch)
	}(ch)

	go func(ch <-chan model.OrderStatus) {
		orderStatus, isChanelOpen := <-ch
		if isChanelOpen {
			createdOrder.Status = orderStatus
			createdOrder, _ = service.orderRepository.UpdateOrder(createdOrder)
		}
	}(ch)

	return createdOrder, nil
}

func (service *OrderService) initializeAndSaveOrder(orderDto dto.OrderDto) model.Order {
	order := mapper.MapOrderToDomain(orderDto)
	var orderPriceTotal float32
	for _, v := range order.Items {
		pizza, _ := service.pizzaService.GetPizzaByName(v.PizzaName)
		orderPriceTotal += pizza.Price * float32(v.Quantity)
	}
	order.Price = orderPriceTotal
	order.Status = model.IN_PREPARATION

	createdOrder, err := service.orderRepository.CreateOrder(order)
	if err != nil {
		fmt.Println(err)
	}
	return createdOrder
}

func (service *OrderService) CheckOrderStatus(orderId int) (model.OrderStatus, error) {
	orderStatus, err := service.orderRepository.CheckOrderStatus(orderId)
	if err != nil {
		fmt.Println(err)
	}
	return orderStatus, err
}

func (service *OrderService) CancelOrder(orderId int) (model.Order, error) {
	var order model.Order
	if !service.checkIfOrderCanBeCancelled(orderId) {
		return order, fmt.Errorf("error can't be cancelled")
	}

	cancelledOrder, err := service.orderRepository.CancelOrder(orderId)
	if err != nil {
		fmt.Println(err)
	}
	return cancelledOrder, err
}

func (service *OrderService) checkIfOrderCanBeCancelled(orderId int) bool {
	order, _ := service.orderRepository.GetById(orderId)

	if order.Status == model.READY_TO_BE_DELIVERED || order.Status == model.CANCELLED {
		return false
	}
	return true
}

func (service *OrderService) CancelOrderRegardlessStatus(orderId int) (model.Order, error) {
	cancelledOrder, err := service.orderRepository.CancelOrder(orderId)
	if err != nil {
		fmt.Println(err)
	}
	return cancelledOrder, err
}