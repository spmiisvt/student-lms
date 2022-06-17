package handlers

import (
	"github.com/gin-gonic/gin"
	"student-lms/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth", h.signUp)
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in", h.signIn)
	}
	return router
}
