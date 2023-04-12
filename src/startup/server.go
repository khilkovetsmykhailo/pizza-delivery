package startup

import (
	gin "github.com/gin-gonic/gin"
	"github.com/marijakljestan/golang-web-app/src/api"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
	service "github.com/marijakljestan/golang-web-app/src/domain/service"
	"github.com/marijakljestan/golang-web-app/src/infrastructure/persistence"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Start() {
	pizzaRepository := server.initPizzaRepository()
	pizzaService := server.initPizzaService(pizzaRepository)
	pizzaHandler := api.NewPizzaController(pizzaService)

	orderRepository := server.initOrderRepository()
	orderService := server.initOrderService(orderRepository, pizzaService)
	orderHandler := api.NewOrderController(orderService)

	userRepository := server.initUserRepository()
	userService := server.initUserService(userRepository)
	userHandler := api.NewUserController(userService)

	router := gin.Default()
	router.GET("/pizza", pizzaHandler.GetMenu)
	router.POST("/pizza", pizzaHandler.AddPizzaToMenu)
	router.DELETE("/pizza/:name", pizzaHandler.DeletePizzaFromMenu)

	router.POST("/order", orderHandler.CreateOrder)
	router.GET("/order/:id", orderHandler.CheckOrderStatus)
	router.PUT("/order/:id", orderHandler.CancelOrder)

	router.POST("/user/register", userHandler.RegisterUser)
	router.Run("localhost:8080")
}

func (server *Server) initPizzaRepository() repository.PizzaRepository {
	return persistence.NewOrderInMemoryRepository()
}

func (server *Server) initPizzaService(orderRepository repository.PizzaRepository) *service.PizzaService {
	return service.NewPizzaService(orderRepository)
}

func (server *Server) initOrderRepository() repository.OrderRepository {
	return persistence.NewOrderInmemoryRepository()
}

func (server *Server) initOrderService(orderRepository repository.OrderRepository, pizzaService *service.PizzaService) *service.OrderService {
	return service.NewOrderService(orderRepository, pizzaService)
}

func (server *Server) initUserRepository() repository.UserRepository {
	return persistence.NewUserInmemoryRepository()
}

func (server *Server) initUserService(userRepository repository.UserRepository) *service.UserService {
	return service.NewUserService(userRepository)
}
