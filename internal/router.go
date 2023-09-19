package internal

import (
	"github.com/gin-gonic/gin"
	"medods_task/internal/app/handlers"
)

type RouteHandler struct {
	*handlers.AuthHandler
	*handlers.MainHandler
}

func NewRouteHandler() *RouteHandler {
	return &RouteHandler{
		AuthHandler: handlers.NewAuthHandler(),
		MainHandler: handlers.NewMainHandler(),
	}
}

func (r *RouteHandler) Handle(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.GET("/login", r.Login)
		auth.POST("/refresh", r.Refresh)
	}
}
