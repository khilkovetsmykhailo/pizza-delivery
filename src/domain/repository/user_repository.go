package domain

import domain "github.com/marijakljestan/golang-web-app/src/domain/model"

type UserRepository interface {
	Save(user *domain.User) (string, error)
	GetByUsername(username string) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	DeleteAll()
}
