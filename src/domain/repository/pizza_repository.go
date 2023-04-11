package domain

import domain "github.com/marijakljestan/golang-web-app/src/domain/model"

type PizzaRepository interface {
	GetMenu() ([]domain.Pizza, error)
	AddPizzaToMenu(pizza domain.Pizza) ([]domain.Pizza, error)
	DeletePizzaFromMenu(pizzaName string) ([]domain.Pizza, error)
	FindPizzaByName(name string) (domain.Pizza, error)
}
