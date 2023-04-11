package persistence

import (
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
)

type UserInmemoryRepository struct{}

var users = []domain.User{
	{
		Username: "admin",
		Password: "admin",
		Type:     domain.ADMIN,
	},
	{
		Username: "customer",
		Password: "customer",
		Type:     domain.CUSTOMER,
	},
}

var userCounter int = 0

func NewUserInmemoryRepository() repository.UserRepository {
	return &UserInmemoryRepository{}
}

func (repository *UserInmemoryRepository) Save(user domain.User) (string, error) {
	userCounter++
	user.Id = userCounter
	users = append(users, user)
	return user.Username, nil
}

func (repository UserInmemoryRepository) GetByUsername(username string) (domain.User, error) {
	var user domain.User
	for _, v := range users {
		if v.Username == username {
			user = v
			break
		}
	}
	return user, nil
}
