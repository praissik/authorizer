package database

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoClient() *mongo.Client {
	// Create a new client and connect to the server

	clientOpts := options.Client().ApplyURI(viper.GetString("mongo"))
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		panic(err)
	}
	//defer func() {
	//	if err = client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil { // ping not working, probably
		panic(err)
	}
	return client
}
