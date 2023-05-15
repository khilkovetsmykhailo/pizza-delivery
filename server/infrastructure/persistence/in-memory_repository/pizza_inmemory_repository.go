package in_memory_repository

import (
	"errors"
	domain "github.com/marijakljestan/golang-web-app/server/domain/model"
	repository "github.com/marijakljestan/golang-web-app/server/domain/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PizzaInmemoryRepository struct{}

var pizzaMenu = []*domain.Pizza{
	{
		Id:          getObjectId("723b0cc3a34d25d8567f9f84"),
		Name:        "Margarita",
		Description: "Margarita description",
		Price:       650.50,
	},
	{
		Id:          getObjectId("723b0cc3a34d25d8567f9f85"),
		Name:        "Capricciosa",
		Description: "Capricciosa description",
		Price:       750.50,
	},
}

func NewOrderInMemoryRepository() repository.PizzaRepository {
	return &PizzaInmemoryRepository{}
}

func (repository *PizzaInmemoryRepository) GetAll() ([]*domain.Pizza, error) {
	return pizzaMenu, nil
}

func (repository *PizzaInmemoryRepository) Insert(pizza *domain.Pizza) ([]*domain.Pizza, error) {
	pizza.Id = primitive.NewObjectID()
	pizzaMenu = append(pizzaMenu, pizza)
	return pizzaMenu, nil
}

func (repository *PizzaInmemoryRepository) Delete(pizzaName string) ([]*domain.Pizza, error) {
	for i, v := range pizzaMenu {
		if v.Name == pizzaName {
			pizzaMenu = append(pizzaMenu[:i], pizzaMenu[i+1:]...)
			return pizzaMenu, nil
		}
	}
	return pizzaMenu, errors.New("pizza with provided name does not exist")
}

func (repository *PizzaInmemoryRepository) GetPizzaByName(pizzaName string) (*domain.Pizza, error) {
	var pizza domain.Pizza
	for _, v := range pizzaMenu {
		if v.Name == pizzaName {
			pizza = *v
			return &pizza, nil
		}
	}
	return &pizza, errors.New("pizza with provided name does not exist")
}

func (repository *PizzaInmemoryRepository) DeleteAll() {
	pizzaMenu = []*domain.Pizza{}
}
