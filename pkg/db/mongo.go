package db

import (
	errors "authorizer/pkg/error"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

const (
	CollectionAccounts          = "accounts"
	CollectionRegistersRequests = "registers_requests"
)

func GetMongoClient() (*mongo.Client, func(), error) {
	// Create a new client and connect to the server

	clientOpts := options.Client().ApplyURI(viper.GetString("mongo.url"))
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Println(err.Error())
		return nil, nil, fmt.Errorf(errors.SomethingWentWrong)
	}

	// Ping the primary
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client, func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Println(err.Error())
		}
	}, nil
}
