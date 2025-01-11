package mongo

import (
	"context"
	"fmt"
	cim "github.com/LucxLab/cim-service/internal"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const host = "localhost"
const databaseName = "cim"

type Database struct {
	client   *mongo.Client
	database *mongo.Database
}

func (d *Database) InsertOne(collectionName string, object interface{}) (identifier string, err error) {
	collection := d.database.Collection(collectionName)
	result, err := collection.InsertOne(context.TODO(), object)
	if err != nil {
		return "", err
	}

	objectId := result.InsertedID.(bson.ObjectID)
	return objectId.Hex(), nil
}

func (d *Database) Close() error {
	if err := d.client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}

func NewDatabase() cim.Database {
	opt := options.Client().ApplyURI("mongodb://" + host)

	client, err := mongo.Connect(opt)
	if err != nil {
		panic(err) //TODO: handle error
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err) //TODO: handle error
	}

	database := client.Database(databaseName)
	fmt.Println("Successfully connected to MongoDB: ", host)
	return &Database{
		client:   client,
		database: database,
	}
}
