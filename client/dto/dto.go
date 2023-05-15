package dto

type Pizza struct {
	Name        string  `json:"name",validate:"required,max=32"`
	Description string  `json:"description"`
	Price       float64 `json:"price",validate:"required,numeric"`
}

type Order struct {
	Id               string      `json:"id"`
	CustomerUsername string      `json:"customer_username"`
	Status           string      `json:"status"`
	Price            float64     `json:"price"`
	Items            []OrderItem `json:"items"`
}

type OrderItem struct {
	PizzaName string `json:"pizza_name",validate:"required"`
	Quantity  int    `json:"quantity",validate:"required,numeric"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type MenuResponse struct {
	Menu []Pizza `json:"menu"`
}

type CreateOrderResponse struct {
	Order Order `json:"order"`
}

type GetOrderStatusResponse struct {
	OrderStatus string `json:"status"`
}

type CancelOrderResponse struct {
	Order Order `json:"order"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
