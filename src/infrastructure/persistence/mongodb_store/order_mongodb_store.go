package mongodb_store

import (
	"context"
	"fmt"
	domain "github.com/marijakljestan/golang-web-app/src/domain/model"
	repository "github.com/marijakljestan/golang-web-app/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderMongoDBStore struct {
	orderCollection *mongo.Collection
}

func NewOrderMongoDBStore(client *mongo.Client) repository.OrderRepository {
	orderCollection := client.Database("pizza-delivery-app").Collection("orders")
	return &OrderMongoDBStore{
		orderCollection: orderCollection,
	}
}

func (store *OrderMongoDBStore) Save(order domain.Order) (domain.Order, error) {
	order.Id = primitive.NewObjectID()
	_, err := store.orderCollection.InsertOne(context.TODO(), order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (store *OrderMongoDBStore) CheckOrderStatus(orderId primitive.ObjectID) (domain.OrderStatus, error) {
	order, err := store.GetById(orderId)
	return order.Status, err
}

func (store *OrderMongoDBStore) CancelOrder(orderId primitive.ObjectID) (*domain.Order, error) {
	_, err := store.orderCollection.UpdateOne(context.TODO(),
		bson.M{"_id": orderId},
		bson.D{{"$set", bson.D{{"status", domain.CANCELLED}}}},
	)
	if err != nil {
		fmt.Println(err)
	}
	return store.GetById(orderId)
}

func (store *OrderMongoDBStore) GetById(orderId primitive.ObjectID) (*domain.Order, error) {
	filter := bson.M{"_id": orderId}
	return store.findOne(filter)
}

func (store *OrderMongoDBStore) Update(order domain.Order) (*domain.Order, error) {
	_, err := store.orderCollection.UpdateOne(context.TODO(),
		bson.M{"_id": order.Id},
		bson.D{{"$set", bson.D{{"status", order.Status}}}},
	)
	if err != nil {
		fmt.Println(err)
	}
	return store.GetById(order.Id)
}

func (store *OrderMongoDBStore) findOne(filter interface{}) (order *domain.Order, err error) {
	result := store.orderCollection.FindOne(context.TODO(), filter)
	err = result.Decode(&order)
	return
}

func (store *OrderMongoDBStore) findAll(filter interface{}) (orders []domain.Order, err error) {
	cursor, err := store.orderCollection.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	return decodeOrder(cursor)
}

func decodeOrder(cursor *mongo.Cursor) (orders []domain.Order, err error) {
	for cursor.Next(context.TODO()) {
		var order domain.Order
		err = cursor.Decode(order)
		if err != nil {
			return
		}
		orders = append(orders, order)
	}
	err = cursor.Err()
	return
}
