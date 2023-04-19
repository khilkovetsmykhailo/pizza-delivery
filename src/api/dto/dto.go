package dto

type PizzaDto struct {
	Name        string  `json:"name",validate:"required,max=32"`
	Description string  `json:"description"`
	Price       float32 `json:"price",validate:"required,numeric"`
}

type OrderDto struct {
	CustomerUsername string         `json:"customer_username"`
	Price            float32        `json:"price"`
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
