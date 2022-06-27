package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func NewConnectMongo() {
	clientOptions := options.Client().ApplyURI(MongoConn())
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	MongoDB = client.Database("one-click-test")
	// return client
}

/* func GetCollection(collectionName string) *mongo.Collection {
	// fmt.Println(MongoDB)
	return MongoDB.Database("one-click-test").Collection(collectionName)
} */

/* func InitDB() {
	MongoDB = NewConnectMongo()
}
*/
