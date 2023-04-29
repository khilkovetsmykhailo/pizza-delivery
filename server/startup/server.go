package startup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marijakljestan/golang-web-app/server/api"
	repository "github.com/marijakljestan/golang-web-app/server/domain/repository"
	"github.com/marijakljestan/golang-web-app/server/domain/service"
	"github.com/marijakljestan/golang-web-app/server/infrastructure/persistence/in-memory_repository"
	"github.com/marijakljestan/golang-web-app/server/infrastructure/persistence/mongodb_store"
	"github.com/marijakljestan/golang-web-app/server/middleware"
	config2 "github.com/marijakljestan/golang-web-app/server/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
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
	mongoClient := server.initMongoClient()

	//pizzaRepository := server.initPizzaRepository()
	pizzaRepository := server.initPizzaMongoDBStore(mongoClient)
	pizzaService := server.initPizzaService(pizzaRepository)
	pizzaHandler := api.NewPizzaController(pizzaService)

	//orderRepository := server.initOrderRepository()
	orderRepository := server.initOrderMongoDBStore(mongoClient)
	orderService := server.initOrderService(orderRepository, pizzaService)
	orderHandler := api.NewOrderController(orderService)

	//userRepository := server.initUserRepository()
	userRepository := server.initUserMongoDBStore(mongoClient)
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
		orderRoutes.POST("", middleware.AuthorizeJWT("CUSTOMER"), orderHandler.CreateOrder)
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

	addr := fmt.Sprintf(":%s", server.config.Port)
	err := router.Run(addr)
	if err != nil {
		fmt.Println("Error! Cannot start server.")
	}
}

func (server *Server) initPizzaRepository() repository.PizzaRepository {
	return in_memory_repository.NewOrderInMemoryRepository()
}

func (server *Server) initPizzaMongoDBStore(client *mongo.Client) repository.PizzaRepository {
	store := mongodb_store.NewPizzaMongoDBStore(client)
	store.DeleteAll()
	for _, pizza := range pizzaMenu {
		_, err := store.Insert(pizza)
		if err != nil {
			fmt.Println(err)
		}
	}
	return store
}

func (server *Server) initPizzaService(orderRepository repository.PizzaRepository) *service.PizzaService {
	return service.NewPizzaService(orderRepository)
}

func (server *Server) initOrderRepository() repository.OrderRepository {
	return in_memory_repository.NewOrderInmemoryRepository()
}

func (server *Server) initOrderMongoDBStore(client *mongo.Client) repository.OrderRepository {
	store := mongodb_store.NewOrderMongoDBStore(client)
	store.DeleteAll()
	return store
}

func (server *Server) initOrderService(orderRepository repository.OrderRepository, pizzaService *service.PizzaService) *service.OrderService {
	return service.NewOrderService(orderRepository, pizzaService)
}

func (server *Server) initUserRepository() repository.UserRepository {
	return in_memory_repository.NewUserInmemoryRepository()
}

func (server *Server) initUserMongoDBStore(client *mongo.Client) repository.UserRepository {
	store := mongodb_store.NewUsersMongoDBStore(client)
	store.DeleteAll()
	for _, user := range Users {
		_, err := store.Save(user)
		if err != nil {
			fmt.Println(err)
		}
	}
	return store
}

func (server *Server) initUserService(userRepository repository.UserRepository) *service.UserService {
	return service.NewUserService(userRepository)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := mongodb_store.GetClient(server.config.DBHost, server.config.DBPort)
	if err != nil {
		fmt.Println(err)
	}
	return client
}
