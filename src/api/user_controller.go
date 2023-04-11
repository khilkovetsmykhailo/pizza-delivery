package api

import (
	"github.com/gin-gonic/gin"
	"github.com/marijakljestan/golang-web-app/src/api/dto"
	"github.com/marijakljestan/golang-web-app/src/domain/service"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (handler *UserController) RegisterUser(ctx *gin.Context) {
	var userDto dto.UserDto
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided!"})
		return
	}
	username, err := handler.userService.RegisterCustomer(userDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving customer!"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"user": username})
}
