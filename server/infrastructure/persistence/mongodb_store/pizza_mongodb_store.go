package mongodb_store

import (
	"context"
	"fmt"
	model "github.com/marijakljestan/golang-web-app/server/domain/model"
	repository "github.com/marijakljestan/golang-web-app/server/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PizzaMongoDBStore struct {
	pizzaCollection *mongo.Collection
}

func NewPizzaMongoDBStore(client *mongo.Client) repository.PizzaRepository {
	pizzaCollection := client.Database("pizza-delivery-app").Collection("pizza")
	return &PizzaMongoDBStore{
		pizzaCollection: pizzaCollection,
	}
}

func (store *PizzaMongoDBStore) GetAll() ([]*model.Pizza, error) {
	filter := bson.D{{}}
	return store.filterAll(filter)
}

func (store *PizzaMongoDBStore) Insert(pizza *model.Pizza) ([]*model.Pizza, error) {
	pizza.Id = primitive.NewObjectID()
	_, err := store.pizzaCollection.InsertOne(context.TODO(), pizza)
	if err != nil {
		fmt.Println(err)
	}
	return store.GetAll()
}

func (store *PizzaMongoDBStore) Delete(pizzaName string) ([]*model.Pizza, error) {
	filter := bson.M{"name": pizzaName}
	_, err := store.pizzaCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return store.GetAll()
}

func (store *PizzaMongoDBStore) GetPizzaByName(name string) (*model.Pizza, error) {
	filter := bson.M{"name": name}
	return store.filterOne(filter)
}

func (store *PizzaMongoDBStore) DeleteAll() {
	store.pizzaCollection.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *PizzaMongoDBStore) filterOne(filter interface{}) (pizza *model.Pizza, err error) {
	result := store.pizzaCollection.FindOne(context.TODO(), filter)
	err = result.Decode(&pizza)
	return
}

func (store *PizzaMongoDBStore) filterAll(filter interface{}) ([]*model.Pizza, error) {
	cursor, err := store.pizzaCollection.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodePizza(cursor)
}

func decodePizza(cursor *mongo.Cursor) (pizzaMenu []*model.Pizza, err error) {
	for cursor.Next(context.TODO()) {
		var pizza model.Pizza
		err = cursor.Decode(&pizza)
		if err != nil {
			return
		}
		pizzaMenu = append(pizzaMenu, &pizza)
	}
	err = cursor.Err()
	return
}
