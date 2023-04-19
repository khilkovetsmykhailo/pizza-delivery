package startup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marijakljestan/golang-web-app/src/api"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
	"github.com/marijakljestan/golang-web-app/src/domain/service"
	"github.com/marijakljestan/golang-web-app/src/infrastructure/persistence"
	"github.com/marijakljestan/golang-web-app/src/middleware"
	config2 "github.com/marijakljestan/golang-web-app/src/startup/config"
)

type Server struct {
	config *config2.Config
}

func NewServer(config *config2.Config) *Server {
	return &Server{
		config: config,
	}
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

	pizzaRoutes := router.Group("/pizza")
	{
		pizzaRoutes.GET("", pizzaHandler.GetMenu)
		pizzaRoutes.POST("", middleware.AuthorizeJWT("ADMIN"), pizzaHandler.AddPizzaToMenu)
		pizzaRoutes.DELETE("/:name", middleware.AuthorizeJWT("ADMIN"), pizzaHandler.DeletePizzaFromMenu)
	}

	orderRoutes := router.Group("/order")
	{
		orderRoutes.POST("", orderHandler.CreateOrder)
		orderRoutes.GET("/status/:id", orderHandler.CheckOrderStatus)
		orderRoutes.PUT("/cancel/:id", orderHandler.CancelOrder)
		orderRoutes.PUT("/:id", middleware.AuthorizeJWT("ADMIN"), orderHandler.CancelOrderRegardlessStatus)
	}

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("", userHandler.GetAll)
		userRoutes.POST("/register", userHandler.RegisterUser)
		userRoutes.POST("/login", userHandler.Login)
	}

	addr := fmt.Sprintf("%s:%s", server.config.Host, server.config.Port)
	err := router.Run(addr)
	if err != nil {
		fmt.Println("Error! Cannot start server.")
	}
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
