package dto

type PizzaDto struct {
	Name        string  `json:"name",validate:"required,max=32"`
	Description string  `json:"description"`
	Price       float64 `json:"price",validate:"required,numeric"`
}

type OrderDto struct {
	Id               string         `json:"id"`
	CustomerUsername string         `json:"customer_username"`
	Price            float64        `json:"price"`
	Status           string         `json:"status"`
	Items            []OrderItemDto `json:"items"`
}

type OrderItemDto struct {
	PizzaName string `json:"pizza_name",validate:"required,"`
	Quantity  int    `json:"quantity",validate:"required,numeric"`
}

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
