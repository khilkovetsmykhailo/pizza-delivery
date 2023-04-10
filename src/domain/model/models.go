package domain

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	UserType UserType `json:"user_type"`
}

type Pizza struct {
	Id          int     `json:"id""`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type Order struct {
	Id       int         `json:"id"`
	Customer User        `json:"customer"`
	Status   OrderStatus `json:"status"`
	Price    float32     `json:"price"`
	Items    []OrderItem `json:"items"`
}

type OrderItem struct {
	Pizza    Pizza `json:"pizza"`
	Quantity int   `json:"quantity"`
}
