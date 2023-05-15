package mongodb_store

import (
	"context"
	model "github.com/marijakljestan/golang-web-app/server/domain/model"
	repository "github.com/marijakljestan/golang-web-app/server/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUsersMongoDBStore(client *mongo.Client) repository.UserRepository {
	userCollection := client.Database("pizza-delivery-app").Collection("users")
	return &UserMongoDBStore{
		users: userCollection,
	}
}

func (store *UserMongoDBStore) Save(user *model.User) (string, error) {
	user.Id = primitive.NewObjectID()
	_, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

func (store *UserMongoDBStore) GetByUsername(username string) (*model.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll() ([]*model.User, error) {
	filter := bson.D{{}}
	return store.filterAll(filter)
}

func (store *UserMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (user *model.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func (store *UserMongoDBStore) filterAll(filter interface{}) ([]*model.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeUser(cursor)
}

func decodeUser(cursor *mongo.Cursor) (users []*model.User, err error) {
	for cursor.Next(context.TODO()) {
		var user model.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}
