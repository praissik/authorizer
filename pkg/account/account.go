package account

import (
	"authorizer/pkg/db"
	errors "authorizer/pkg/error"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Entity struct {
	//ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email         string `bson:"email" json:"email,omitempty"`
	Password      string `bson:"password" json:"password,omitempty"`
	LastLoginDate int64  `bson:"last_login_date" json:"last_login_date,omitempty"`
	CreateDate    int64  `bson:"create_date" json:"create_date,omitempty"`
}

func (e *Entity) Get() {

	return
}

func (e *Entity) Create() (*mongo.InsertOneResult, error) {
	mongoClient, deferF, err := db.GetMongoClient()
	defer deferF()
	if err != nil {
		return nil, err
	}

	res, err := mongoClient.
		Database(viper.GetString("mongo.db")).
		Collection(db.CollectionAccounts).
		InsertOne(nil, e)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf(errors.SomethingWentWrong)
	}
	return res, nil
}

func (e *Entity) Update() interface{} {

	return nil
}

func (e *Entity) Delete() interface{} {

	return nil
}

func GetAccountByEmail(email string) (*Entity, error) {
	var accountEntity *Entity
	mongoClient, deferF, err := db.GetMongoClient()
	defer deferF()
	if err != nil {
		return nil, err
	}

	err = mongoClient.
		Database(viper.GetString("mongo.db")).
		Collection(db.CollectionAccounts).
		FindOne(context.TODO(), bson.M{"email": email}).
		Decode(&accountEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, fmt.Errorf(errors.SomethingWentWrong)
	}
	return accountEntity, nil
}
