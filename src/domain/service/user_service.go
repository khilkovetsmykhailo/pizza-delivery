package service

import (
	"fmt"
	"github.com/marijakljestan/golang-web-app/src/api/dto"
	"github.com/marijakljestan/golang-web-app/src/domain/mapper"
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
	utils "github.com/marijakljestan/golang-web-app/src/util"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) RegisterCustomer(userDto dto.UserDto) (string, error) {
	user := mapper.MapUserToDomain(userDto)
	user.Role = domain.CUSTOMER
	user.Password = utils.HashPassword(user.Password)
	username, err := service.userRepository.Save(user)
	if err != nil {
		fmt.Println(err)
	}
	return username, err
}

func (service *UserService) GetByUsername(username string) (domain.User, error) {
	user, err := service.userRepository.GetByUsername(username)
	if err != nil {
		fmt.Println(err)
	}
	return user, err
}
