package middlewares

import (
	"github.com/gin-gonic/gin"
	"medods_task/internal/app/handlers/responses"
)

func InitializerMiddleware() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		responses.GinContext = ginContext

		ginContext.Next()
	}
}
