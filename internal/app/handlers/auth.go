package handlers

import (
	"fmt"
	"github.com/beevik/guid"
	"github.com/gin-gonic/gin"
	"medods_task/internal/app/actions"
	"medods_task/internal/app/handlers/responses"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) Login(ginContext *gin.Context) {
	guidParam := ginContext.Query("guid")
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

	responses.Done(result)

}

func (a *AuthHandler) Refresh(ginContext *gin.Context) {

}
