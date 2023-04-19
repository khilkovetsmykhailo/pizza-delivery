package dto

type Pizza struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type Order struct {
	Id               int         `json:"id"`
	CustomerUsername string      `json:"customer_username"`
	Status           int         `json:"status"`
	Price            float32     `json:"price"`
	Items            []OrderItem `json:"items"`
}

type OrderItem struct {
	PizzaName string `json:"pizza_name"`
	Quantity  int    `json:"quantity"`
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
