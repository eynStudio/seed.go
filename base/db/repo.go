package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	Name string
}

func (repo MongoRepo) Coll() *mongo.Collection {
	return DefaultMongo.Collection(repo.Name)
}
