package responses

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	GinContext *gin.Context
)

func Error(err error) {
	switch err {
	default:
		log.Println(color.RedString(err.Error()))
		GinContext.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "responses": err.Error()})
	}
}

func Done(result interface{}) {
	if result == nil {
		result = "successfully"
	}

	GinContext.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": result})
}
