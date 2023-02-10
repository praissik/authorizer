package account

import (
	"authorizer/pkg/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	ID       string
	Email    string
	Password string
}

func IsEmailExists(email string) (bool, error) {
	mongoClient := database.GetMongoClient()

	var result Account
	//var result bson.M
	err := mongoClient.
		Database("test").
		Collection("account").
		FindOne(context.TODO(), bson.M{"email": email}).
		Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
