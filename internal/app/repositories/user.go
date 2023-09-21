package repositories

import (
	"context"
	"github.com/beevik/guid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"medods_task/configs"
	"medods_task/internal/app/models"
)

type UserRepository struct {
	*MainRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{MainRepository: NewMainRepository()}
}

func (u *UserRepository) GetByGUID(guid *guid.Guid) (*models.User, error) {
	var targetUser models.User
	err := configs.DbObject.Where("guid = ?", guid).First(&targetUser).Error
	if err != nil {
		return nil, err
	}

	return &targetUser, nil
}

type UserDAO struct {
	*MongoRepository
}

func NewUserDAO() *UserDAO {
	return &UserDAO{
		MongoRepository: NewMongoRepository("users"),
	}
}

func (d *UserDAO) InsertOne(item bson.M) (*mongo.InsertOneResult, error) {
	result, err := d.Collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}

	return result, nil
}
