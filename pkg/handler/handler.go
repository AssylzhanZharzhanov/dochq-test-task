package handler

import (
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		answers := api.Group("/answers")
		{
			answers.POST("", h.createAnswer)
			answers.GET("", h.getAnswer)
			answers.PUT("", h.updateAnswer)
			answers.DELETE("", h.deleteAnswer)
			answers.GET("/history", h.getEvents)
		}
	}

	return router
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
