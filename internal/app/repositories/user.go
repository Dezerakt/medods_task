package repositories

import (
	"github.com/beevik/guid"
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
