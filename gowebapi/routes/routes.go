package routes

import (
	"gowebapi/handlers"

	"github.com/gin-gonic/gin"
)

type TemperatureRoutes struct{}

func (t TemperatureRoutes) ConfigureRoutes(router *gin.Engine) {
	router.GET("/temperatures", handlers.GetTemperatures)
	router.POST("/temperatures", handlers.CreateTemperature)
	router.GET("/temperatures/:id", handlers.GetTemperatureByID)
	router.PUT("/temperatures/:id", handlers.UpdateTemperature)
	router.DELETE("/temperatures/:id", handlers.DeleteTemperature)
}
