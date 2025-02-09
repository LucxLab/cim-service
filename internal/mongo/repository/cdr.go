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

func (c *mongoCdr) CreateFileMetadata(upload *cdr.FileMetadata) error {
	collection := c.database.Cim.Collection(cdrFileMetadata)
	createFileMetadata, err := data.ToCreateFileMetadata(upload)
	if err != nil {
		return err
	}

	insertResult, err := collection.InsertOne(context.TODO(), createFileMetadata)
	if err != nil {
		return err
	}

	objectId := insertResult.InsertedID.(bson.ObjectID)
	upload.Id = objectId.Hex()
	return nil
}

func (c *mongoCdr) UploadSucceeded(id string, fileLocation string) error {
	collection := c.database.Cim.Collection(cdrFileMetadata)
	uploadSucceeded := data.ToUploadSucceeded(fileLocation)

	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updateActions := bson.M{"$set": uploadSucceeded}
	_, err = collection.UpdateByID(context.TODO(), objectId, updateActions)
	if err != nil {
		return err
	}
	return nil
}

func (c *mongoCdr) UploadFailed(id string) error {
	collection := c.database.Cim.Collection(cdrFileMetadata)
	uploadFailed := data.ToUploadFailed()

	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updateActions := bson.M{"$set": uploadFailed}
	_, err = collection.UpdateByID(context.TODO(), objectId, updateActions)
	if err != nil {
		return err
	}
	return nil
}

func NewMongoCdr(database *mongo.Database) cdr.DatabaseRepository {
	return &mongoCdr{database: database}
}
