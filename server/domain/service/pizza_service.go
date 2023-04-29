package service

import (
	"errors"
	"fmt"
	"github.com/marijakljestan/golang-web-app/server/api/dto"
	"github.com/marijakljestan/golang-web-app/server/domain/mapper"
	model "github.com/marijakljestan/golang-web-app/server/domain/model"
	repository "github.com/marijakljestan/golang-web-app/server/domain/repository"
)

type PizzaService struct {
	pizzaRepository repository.PizzaRepository
}

func NewPizzaService(pizzaRepository repository.PizzaRepository) *PizzaService {
	return &PizzaService{
		pizzaRepository: pizzaRepository,
	}
}

func (service *PizzaService) ListMenu() ([]*model.Pizza, error) {
	menu, err := service.pizzaRepository.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	return menu, nil
}

func (service *PizzaService) AddPizzaToMenu(pizzaDto dto.PizzaDto) ([]*model.Pizza, error) {
	pizza := mapper.MapPizzaToDomain(pizzaDto)
	menu, err := service.pizzaRepository.Insert(&pizza)
	if err != nil {
		fmt.Println(err)
	}
	return menu, err
}

func (service *PizzaService) DeletePizzaFromMenu(pizzaName string) ([]*model.Pizza, error) {
	if pizzaExists := service.checkIfPizzaExists(pizzaName); !pizzaExists {
		return nil, errors.New("pizza with provided name does not exist")
	}
	menu, err := service.pizzaRepository.Delete(pizzaName)
	if err != nil {
		fmt.Println(err)
	}
	return menu, nil
}

func (service *PizzaService) GetPizzaByName(pizzaName string) (*model.Pizza, error) {
	pizza, err := service.pizzaRepository.GetPizzaByName(pizzaName)
	if err != nil {
		fmt.Println(err)
	}
	return pizza, nil
}

func (service *PizzaService) checkIfPizzaExists(pizzaName string) bool {
	_, err := service.pizzaRepository.GetPizzaByName(pizzaName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
