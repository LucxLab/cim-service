package repositories

import (
	"context"
	"github.com/LucxLab/cim-service/internal/mongo"
	"github.com/LucxLab/cim-service/internal/mongo/data"
	"github.com/LucxLab/cim-service/internal/repositories"
	"github.com/LucxLab/cim-service/internal/upload"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const cdrFileMetadata = "cdr_files_metadata"

type mongoCdr struct {
	database *mongo.Database
}

func (c *mongoCdr) CreateMetadata(fileCreation *upload.FileCreation) (id string, err error) {
	collection := c.database.Cim.Collection(cdrFileMetadata)

	metadataCreation, adapterErr := data.NewFileMetadataCreation(fileCreation)
	if adapterErr != nil {
		return "", adapterErr
	}

	insertResult, insertErr := collection.InsertOne(context.TODO(), metadataCreation)
	if insertErr != nil {
		return "", insertErr
	}

	objectId := insertResult.InsertedID.(bson.ObjectID)
	return objectId.Hex(), nil
}

func (c *mongoCdr) UploadSucceeded(id string, fileLocation string) error {
	collection := c.database.Cim.Collection(cdrFileMetadata)
	uploadSucceeded := data.NewFileUploadSucceeded(fileLocation)

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
	uploadFailed := data.NewFileUploadFailed()

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

func NewMongoCdr(database *mongo.Database) repositories.CdrDatabase {
	return &mongoCdr{database: database}
}
