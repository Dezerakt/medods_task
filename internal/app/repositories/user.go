package repositories

import (
	context "context"
	"github.com/beevik/guid"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (d *UserDAO) InsertOne(item *models.User) (*models.User, error) {
	item.ID = primitive.NewObjectID()
	_, err := d.Collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (d *UserDAO) Find(query *models.User) (*models.User, error) {
	var targetUser models.User
	err := d.Collection.FindOne(context.Background(), query).Decode(&targetUser)
	if err != nil {
		return nil, err
	}

	return &targetUser, nil
}
