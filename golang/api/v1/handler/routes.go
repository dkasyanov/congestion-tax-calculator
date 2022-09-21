package handler

import (
	"congestion-calculator/service"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(app *gin.Engine, s service.IService) {
	v1 := app.Group("api/v1")
	{
		v1.POST("/calculate", postRequestUrl(s))
	}
}
