package handler

import (
	_ "github.com/Relax-FM/Test_effective_mobile_go_2024/docs"
	"github.com/Relax-FM/Test_effective_mobile_go_2024/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// router.GET("/", h.main)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		music := api.Group("/music")
		{
			music.GET("/", h.getAllMusic) // 
			music.GET("/:id", h.getMusicText) //
			music.DELETE("/:id", h.deleteMusicItem) //
			music.PUT("/:id", h.updateMusicItem) // Готово
			music.POST("/", h.addMusicItem) // Готово
		}
	}

	return router
}