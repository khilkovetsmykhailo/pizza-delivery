package in_memory_repository

import (
	"errors"
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInmemoryRepository struct{}

var users = []*domain.User{
	{
		Id:       getObjectId("723b0cc3a34d25d8567f9f82"),
		Username: "admin",
		Password: "$2a$12$4b5bv2fgn31QQboo8vjq0.w/I7iXAUDagIcCJzkDzkLXL4nFOfHgm", //admin
		Role:     domain.ADMIN,
	},
	{
		Id:       getObjectId("723b0cc3a34d25d8567f9f72"),
		Username: "customer",
		Password: "$2a$12$n.qmZtK5oUGyVS0ixhEncOQCoNKOKfylDkGlfGYWJ4Z7d8LrT5j2q", //customer
		Role:     domain.CUSTOMER,
	},
}

func NewUserInmemoryRepository() repository.UserRepository {
	return &UserInmemoryRepository{}
}

func (repository *UserInmemoryRepository) Save(user *domain.User) (string, error) {
	user.Id = primitive.NewObjectID()
	users = append(users, user)
	return user.Username, nil
}

func (repository *UserInmemoryRepository) GetByUsername(username string) (*domain.User, error) {
	var user *domain.User
	for _, v := range users {
		if v.Username == username {
			user = v
			return user, nil
		}
	}
	return user, errors.New("user with provided username does not exist")
}

func (repository *UserInmemoryRepository) GetAll() ([]*domain.User, error) {
	return users, nil
}

func (repository *UserInmemoryRepository) DeleteAll() {
	users = []*domain.User{}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
