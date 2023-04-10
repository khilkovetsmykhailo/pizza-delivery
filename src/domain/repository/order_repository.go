package domain

import domain "github.com/marijakljestan/golang-web-app/src/domain/model"

type OrderRepository interface {
	GetMenu() ([]domain.Pizza, error)
}
