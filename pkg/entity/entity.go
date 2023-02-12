package entity

import "go.mongodb.org/mongo-driver/mongo"

type Entity interface {
	Get()
	Create() (*mongo.InsertOneResult, error)
	Update() interface{}
	Delete() interface{}
}
