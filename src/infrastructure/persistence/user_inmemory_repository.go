package persistence

import (
	"errors"
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
)

type UserInmemoryRepository struct{}

var users = []domain.User{
	{
		Id:       1,
		Username: "admin",
		Password: "$2a$12$4b5bv2fgn31QQboo8vjq0.w/I7iXAUDagIcCJzkDzkLXL4nFOfHgm", //admin
		Role:     domain.ADMIN,
	},
	{
		Id:       2,
		Username: "customer",
		Password: "$2a$12$n.qmZtK5oUGyVS0ixhEncOQCoNKOKfylDkGlfGYWJ4Z7d8LrT5j2q", //customer
		Role:     domain.CUSTOMER,
	},
}
var userCounter = len(users)

func NewUserInmemoryRepository() repository.UserRepository {
	return &UserInmemoryRepository{}
}

func (repository *UserInmemoryRepository) Save(user domain.User) (string, error) {
	userCounter++
	user.Id = userCounter
	users = append(users, user)
	return user.Username, nil
}

func (repository *UserInmemoryRepository) GetByUsername(username string) (domain.User, error) {
	var user domain.User
	for _, v := range users {
		if v.Username == username {
			user = v
			return user, nil
		}
	}
	return user, errors.New("user with provided username does not exist")
}

func (repository *UserInmemoryRepository) GetAll() ([]domain.User, error) {
	return users, nil
}
