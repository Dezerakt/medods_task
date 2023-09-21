package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"medods_task/internal/app/repositories"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) Login(ginContext *gin.Context) {
	/*guidParam := ginContext.Query("guid")
	if guidParam == "" {
		responses.Error(fmt.Errorf("guid parameter not found"))
		return
	}

	parseString, err := guid.ParseString(guidParam)
	if err != nil {
		responses.Error(err)
		return
	}

	loginAction := actions.NewLoginAction(parseString)
	result, err := loginAction.Execute()
	if err != nil {
		responses.Error(err)
		return
	}

	responses.Done(result)*/

	userDao := repositories.NewUserDAO()
	one, err := userDao.InsertOne(bson.M{
		"something": "bruh",
	})
	dump.P(one)

	if err != nil {
		log.Print(err)
	}

}

func (a *AuthHandler) Refresh(ginContext *gin.Context) {

}
