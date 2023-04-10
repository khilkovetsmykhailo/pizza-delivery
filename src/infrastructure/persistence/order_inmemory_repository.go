package persistence

import (
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
)

type OrderInmemoryRepository struct{}

var pizza_menu = []domain.Pizza{
	{
		Id:          1,
		Name:        "Margarita",
		Description: "Margarita description",
		Price:       650.50,
	},
	{
		Id:          2,
		Name:        "Madjarica",
		Description: "Madjarica description",
		Price:       750.50,
	},
}

func NewOrderInMemoryRepository() repository.OrderRepository {
	return &OrderInmemoryRepository{}
}

func (repository *OrderInmemoryRepository) GetMenu() ([]domain.Pizza, error) {
	return pizza_menu, nil
}
