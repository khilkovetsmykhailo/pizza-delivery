package startup

import (
	domain "github.com/marijakljestan/golang-web-app/server/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Users = []*domain.User{
	{
		Id:       getObjectId("723b0cc3a34d25d8567f9f82"),
		Username: "admin",
		Password: "$2a$12$4b5bv2fgn31QQboo8vjq0.w/I7iXAUDagIcCJzkDzkLXL4nFOfHgm", //admin
		Role:     domain.ADMIN,
	},
	{
		Id:       getObjectId("723b0cc3a34d25d8567f9f72"),
		Username: "customer",
		Password: "$2a$12$n.qmZtK5oUGyVS0ixhEncOQCoNKOKfylDkGlfGYWJ4Z7d8LrT5j2q", //customer
		Role:     domain.CUSTOMER,
	},
}

var pizzaMenu = []*domain.Pizza{
	{
		Id:          getObjectId("723b0cc3a34d25d8567f9f84"),
		Name:        "Margarita",
		Description: "Margarita description",
		Price:       650.50,
	},
	{
		Id:          getObjectId("723b0cc3a34d25d8567f9f85"),
		Name:        "Capricciosa",
		Description: "Capricciosa description",
		Price:       750.50,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
