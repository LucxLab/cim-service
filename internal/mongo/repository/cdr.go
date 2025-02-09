package repository

import (
	"context"
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/mongo"
	"github.com/LucxLab/cim-service/internal/mongo/data"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const cdrFileMetadata = "cdr_files_metadata"

type mongoCdr struct {
	database *mongo.Database
}

func (c *mongoCdr) CreateCdrFileMetadata(upload *cdr.FileMetadata) error {
	collection := c.database.Cim.Collection(cdrFileMetadata)
	createCdrFileMetadata, err := data.ToCreateCdrFileMetadata(upload)
	if err != nil {
		return err
	}

	insertResult, err := collection.InsertOne(context.TODO(), createCdrFileMetadata)
	if err != nil {
		return err
	}

	objectId := insertResult.InsertedID.(bson.ObjectID)
	upload.Id = objectId.Hex()
	return nil
}

func (c *mongoCdr) UpdateCdrFileMetadata(upload *cdr.FileMetadata) error {
	collection := c.database.Cim.Collection(cdrFileMetadata)
	updateCdrFileMetadata := data.ToUpdateCdrFileMetadata(upload)

	objectId, err := bson.ObjectIDFromHex(upload.Id)
	if err != nil {
		return err
	}

	updateActions := bson.M{"$set": updateCdrFileMetadata}
	_, err = collection.UpdateByID(context.TODO(), objectId, updateActions)
	if err != nil {
		return err
	}
	return nil
}

func NewMongoCdr(database *mongo.Database) cdr.DatabaseRepository {
	return &mongoCdr{database: database}
}
