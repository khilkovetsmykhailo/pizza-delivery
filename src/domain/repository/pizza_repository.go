package domain

import domain "github.com/marijakljestan/golang-web-app/src/domain/model"

type PizzaRepository interface {
	GetAll() ([]*domain.Pizza, error)
	Insert(pizza *domain.Pizza) ([]*domain.Pizza, error)
	Delete(pizzaName string) ([]*domain.Pizza, error)
	GetPizzaByName(name string) (*domain.Pizza, error)
	DeleteAll()
}
