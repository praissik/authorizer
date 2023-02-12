package register_request

import (
	"authorizer/pkg/db"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Entity struct {
	//ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	AccountID      primitive.ObjectID `bson:"account_id" json:"account_id,omitempty"`
	Token          string             `bson:"token" json:"token,omitempty"`
	GenerationDate int64              `bson:"generation_date" json:"generation_date,omitempty"`
	ExpirationDate int64              `bson:"expiration_date" json:"expiration_date,omitempty"`
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

	return mongoClient.
		Database(viper.GetString("mongo.db")).
		Collection(db.CollectionRegistersRequests).
		InsertOne(nil, e)
}

func (e *Entity) Update() interface{} {

	return nil
}

func (e *Entity) Delete() interface{} {

	return nil
}
