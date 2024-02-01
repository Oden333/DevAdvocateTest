package handler

import (
	"DATest/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//Конфигурация страницы ответа
	router := gin.Default()

	//Конфигурация рутов
	auth := router.Group("/people")
	{
		// Метод для добавления новых людей в формате
		auth.POST("/add", h.create_user)
		// Метод для получения данных с различными фильтрами и пагинацией
		auth.GET("/page/:page", h.get_all_users)
		// Метод для изменения сущности
		auth.POST("/edit/:id", h.update_user_by_id)
		// Метод для удаления сущности
		auth.DELETE("/delete/:id", h.delete_user_by_id)

	}

	return router
}
