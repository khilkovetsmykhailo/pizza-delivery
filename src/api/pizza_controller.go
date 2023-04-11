package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marijakljestan/golang-web-app/src/api/dto"
	"github.com/marijakljestan/golang-web-app/src/domain/service"
	"net/http"
)

type PizzaController struct {
	pizzaService *service.PizzaService
}

func NewPizzaController(service *service.PizzaService) *PizzaController {
	return &PizzaController{
		pizzaService: service,
	}
}

func (handler *PizzaController) GetMenu(ctx *gin.Context) {
	menu, err := handler.pizzaService.ListMenu()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting menu!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": menu})
}

func (handler *PizzaController) AddPizzaToMenu(ctx *gin.Context) {
	var pizzaDto dto.PizzaDto
	if err := ctx.BindJSON(&pizzaDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid arguments passed!"})
		return
	}

	menu, err := handler.pizzaService.AddPizzaToMenu(pizzaDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while adding pizza on menu"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": menu})
}

func (handler *PizzaController) DeletePizzaFromMenu(ctx *gin.Context) {
	var pizzaName string = ctx.Param("name")
	menu, err := handler.pizzaService.DeletePizzaFromMenu(pizzaName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Pizza not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": menu})
}
