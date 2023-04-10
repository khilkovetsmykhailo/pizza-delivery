package startup

import (
	gin "github.com/gin-gonic/gin"
	"github.com/marijakljestan/golang-web-app/src/controllers"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
	service "github.com/marijakljestan/golang-web-app/src/domain/service"
	"github.com/marijakljestan/golang-web-app/src/infrastructure/persistence"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Start() {
	orderRepository := server.initOrderRepository()
	orderService := server.initOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)

	router := gin.Default()
	router.GET("/order/menu", orderController.GetMenu)
	router.Run("localhost:8080")
}

func (server *Server) initServer() {

}

func (server *Server) initOrderRepository() repository.OrderRepository {
	return persistence.NewOrderInMemoryRepository()
}

func (server *Server) initOrderService(orderRepository repository.OrderRepository) *service.OrderService {
	return service.NewOrderService(orderRepository)
}
