package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id",json:"id"`
	Username string             `bson:"username",json:"username",validate:"required"`
	Password string             `bson:"password",json:"password",validate:"required"`
	Role     UserRole           `bson:"role",json:"role"`
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
