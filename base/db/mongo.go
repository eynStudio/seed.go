package db

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	session *mongo.Client
}

var DefaultMongo = &Mongo{}

func OpenMongo() *Mongo {
	host := viper.GetString("mongo.uri")
	db := viper.GetString("mongo.db")
	username := viper.GetString("mongo.username")
	password := viper.GetString("mongo.password")

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    db,
		Username:      username,
		Password:      password,
	}
	clientOpts := options.Client().ApplyURI(host).SetAuth(credential)

	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	DefaultMongo.session = client
	return DefaultMongo
}

func (db Mongo) Close() {
	err := db.session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func (db Mongo) Collection(collection string) *mongo.Collection {
	database := viper.GetString("mongo.db")

	return DefaultMongo.session.Database(database).Collection(collection)
}
