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
	Id          primitive.ObjectID `bson:"_id",json:"id""`
	Name        string             `bson:"name",json:"name",validate:"required,max=32"`
	Description string             `bson:"description",json:"description"`
	Price       float64            `bson:"price",json:"price",validate:"required,numeric"`
}

type Order struct {
	Id               primitive.ObjectID `bson:"_id",json:"id"`
	CustomerUsername string             `bson:"customer_username",json:"customer_username"`
	Status           OrderStatus        `bson:"status",json:"status"`
	Price            float64            `bson:"price",json:"price"`
	Items            []OrderItem        `bson:"items",json:"items"`
}

type OrderItem struct {
	PizzaName string `bson:"pizza_name",json:"pizza_name",validate:"required"`
	Quantity  int    `bson:"quantity",json:"quantity",validate:"required,numeric"`
}
