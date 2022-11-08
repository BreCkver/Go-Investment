package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Conexion() *mongo.Client {
	//os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb+srv://userTransaction:3F5KYgCkMvJZZmEe@cluster0.uywov.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	return client
}
