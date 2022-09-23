package main

import (
	"congestion-calculator/api/v1/handler"
	"congestion-calculator/config"
	"congestion-calculator/repository/db"
	"congestion-calculator/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	monApp := gin.New()

	conf := config.Load()
	dbRepo := db.New(conf)

	calculatorService := service.New(dbRepo, conf)
	handler.RegisterAppHandlers(app, calculatorService)

	handler.RegisterMonitoringHandlers(monApp)

	go monApp.Run(conf.MonitoringPort)
	if err := app.Run(conf.ApplicationPort); err != nil {
		panic(fmt.Sprintf("Cannot start application: %v", err))
	}
}
