package actions

import (
	"errors"
	"github.com/beevik/guid"
	"gorm.io/gorm"
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
	targetUser, err := l.userInit()
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := utils.GenerateTokenCouple(l.GUID)
	if err != nil {
		return nil, err
	}

	l.saveTokenCouple(accessToken, refreshToken, targetUser.ID)

	// Не очень понял из задания - refresh токен в формате в base64 передается только
	// при передаче в запросе, или ещё и при генерации?
	return map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (l *Login) userInit() (*models.User, error) {
	userRepository := repositories.NewUserRepository()

	targetUser, err := userRepository.GetByGUID(l.GUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		createdUser, err2 := userRepository.Create(&models.User{
			GUID: l.GUID.String(),
		})
		if err2 != nil {
			return nil, err2
		}

		targetUser = createdUser.(*models.User)
	} else if err != nil {
		return nil, err
	}

	return targetUser, nil
}

func (l *Login) saveTokenCouple(accessToken string, refreshToken string, userId uint) {

}
