package repository

import (
	"context"
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/mongo"
	"github.com/LucxLab/cim-service/internal/mongo/data"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const cdrUploadCollection = "cdr_files"

type mongoCdr struct {
	database *mongo.Database
}

func (c *mongoCdr) CreateUpload(upload *cdr.Upload) error {
	collection := c.database.Cim.Collection(cdrUploadCollection)
	createUploadData, err := data.ToCreateUpload(upload)
	if err != nil {
		return err
	}

	insertResult, err := collection.InsertOne(context.TODO(), createUploadData)
	if err != nil {
		return err
	}

	objectId := insertResult.InsertedID.(bson.ObjectID)
	upload.Id = objectId.Hex()
	return nil
}

func (c *mongoCdr) UpdateUpload(upload *cdr.Upload) error {
	collection := c.database.Cim.Collection(cdrUploadCollection)
	updateUploadData := data.ToUpdateUpload(upload)

	objectId, err := bson.ObjectIDFromHex(upload.Id)
	if err != nil {
		return err
	}

	updateAction := bson.M{"$set": updateUploadData}
	_, err = collection.UpdateByID(context.TODO(), objectId, updateAction)
	if err != nil {
		return err
	}
	return nil
}

func NewMongoCDR(database *mongo.Database) cdr.DatabaseRepository {
	return &mongoCdr{database: database}
}
