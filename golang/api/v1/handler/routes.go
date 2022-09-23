package handler

import (
	"congestion-calculator/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAppHandlers(app *gin.Engine, s service.IService) {
	v1 := app.Group("api/v1")
	{
		v1.POST("/calculate", postRequestUrl(s))
	}
}

func RegisterMonitoringHandlers(app *gin.Engine) {
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
