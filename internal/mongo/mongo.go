package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const defaultHost = "localhost"
const cimDatabase = "cim"

type Database struct {
	client *mongo.Client
	Cim    *mongo.Database
}

func (m *Database) Close() error {
	return m.client.Disconnect(context.TODO())
}

func NewDatabase() *Database {
	opt := options.Client().ApplyURI("mongodb://" + defaultHost)

	client, err := mongo.Connect(opt)
	if err != nil {
		panic(err) //TODO: handle error
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err) //TODO: handle error
	}
	return &Database{
		client: client,
		Cim:    client.Database(cimDatabase),
	}
}
