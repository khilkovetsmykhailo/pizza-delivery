package api

import (
	"github.com/gin-gonic/gin"
	"github.com/marijakljestan/golang-web-app/src/api/dto"
	"github.com/marijakljestan/golang-web-app/src/domain/service"
	utils "github.com/marijakljestan/golang-web-app/src/util"
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

func (handler *UserController) Login(ctx *gin.Context) {
	var credentials dto.UserDto
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := handler.userService.GetByUsername(credentials.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "User not found"})
		return
	}

	if passwordMatches := utils.ComparePassword(user.Password, credentials.Password); passwordMatches {
		token := utils.GenerateToken(user.Username, user.Role.String())
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully logged in", "token": token})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid credentials"})
}
