package in_memory_repository

import (
	"errors"
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
)

type PizzaInmemoryRepository struct{}

var pizzaMenu = []domain.Pizza{
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
var pizzaIdIncrementer = len(pizzaMenu)

func NewOrderInMemoryRepository() repository.PizzaRepository {
	return &PizzaInmemoryRepository{}
}

func (repository *PizzaInmemoryRepository) GetMenu() ([]domain.Pizza, error) {
	return pizzaMenu, nil
}

func (repository *PizzaInmemoryRepository) AddPizzaToMenu(pizza domain.Pizza) ([]domain.Pizza, error) {
	pizzaIdIncrementer++
	pizza.Id = pizzaIdIncrementer
	pizzaMenu = append(pizzaMenu, pizza)
	return pizzaMenu, nil
}

func (repository *PizzaInmemoryRepository) DeletePizzaFromMenu(pizzaName string) ([]domain.Pizza, error) {
	for i, v := range pizzaMenu {
		if v.Name == pizzaName {
			pizzaMenu = append(pizzaMenu[:i], pizzaMenu[i+1:]...)
			return pizzaMenu, nil
		}
	}
	return pizzaMenu, errors.New("pizza with provided name does not exist")
}

func (repository *PizzaInmemoryRepository) FindPizzaByName(pizzaName string) (domain.Pizza, error) {
	var pizza domain.Pizza
	for _, v := range pizzaMenu {
		if v.Name == pizzaName {
			pizza = v
			return pizza, nil
		}
	}
	return pizza, errors.New("pizza with provided name does not exist")
}
