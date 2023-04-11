package dto

type PizzaDto struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type OrderDto struct {
	CustomerUsername string         `json:"customer_username"`
	Price            float32        `json:"price"`
	Items            []OrderItemDto `json:"items"`
}

type OrderItemDto struct {
	PizzaName string `json:"pizza_name"`
	Quantity  int    `json:"quantity"`
}

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
