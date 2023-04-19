package domain

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username",validate:"required"`
	Password string   `json:"password",validate:"required"`
	Role     UserRole `json:"role"`
}

type Pizza struct {
	Id          int     `json:"id""`
	Name        string  `json:"name",validate:"required,max=32"`
	Description string  `json:"description"`
	Price       float32 `json:"price",validate:"required,numeric"`
}

type Order struct {
	Id               int         `json:"id"`
	CustomerUsername string      `json:"customer_username"`
	Status           OrderStatus `json:"status"`
	Price            float32     `json:"price"`
	Items            []OrderItem `json:"items"`
}

type OrderItem struct {
	Id        int    `json:"id"`
	PizzaName string `json:"pizza_name",validate:"required"`
	Quantity  int    `json:"quantity",validate:"required,numeric"`
}
