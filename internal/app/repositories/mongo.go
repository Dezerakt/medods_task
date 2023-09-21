package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"medods_task/configs"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(collection string) *MongoRepository {
	return &MongoRepository{
		Collection: configs.MongoObject.Database(configs.EnvConfigObject.MongoDb).Collection(collection),
	}
}
