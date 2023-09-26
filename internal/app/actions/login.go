package actions

import (
	"errors"
	"github.com/beevik/guid"
	"github.com/gookit/goutil/dump"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"medods_task/internal/app/models"
	"medods_task/internal/app/repositories"
	"medods_task/internal/app/utils"
)

type Login struct {
	GUID *guid.Guid
}

func NewLoginAction(GUID *guid.Guid) *Login {
	return &Login{GUID: GUID}
}

func (l *Login) Execute() (map[string]interface{}, error) {
	targetUserId, err := l.userInit()
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := utils.GenerateTokenCouple(l.GUID)
	if err != nil {
		return nil, err
	}

	dump.P(targetUserId)

	//l.saveTokenCouple(accessToken, refreshToken, targetUserId)

	// Не очень понял из задания - refresh токен в формате в base64 передается только
	// при передаче в запросе, или ещё и при генерации?
	return map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (l *Login) userInit() (interface{}, error) {
	userDao := repositories.NewUserDAO()

	targetUser, err := userDao.Find(&models.User{
		GUID: l.GUID.String(),
	})

	if errors.Is(err, mongo.ErrNoDocuments) {
		targetUser, err := userDao.InsertOne(&models.User{GUID: l.GUID.String()})
		if err != nil {
			return "", err
		}

		return targetUser.InsertedID, nil
	} else if err != nil {
		return "", err
	}

	return targetUser.ID, nil
}

func (l *Login) saveTokenCouple(accessToken string, refreshToken string, userId primitive.ObjectID) {

}
