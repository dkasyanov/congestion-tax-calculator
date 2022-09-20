package main

import (
	"congestion-calculator/calculator"
	"congestion-calculator/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	VehicleType string   `json:"vehicle_type"`
	Records     []string `json:"records"`
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/calculate", func(ctx *gin.Context) {
		var data RequestBody
		ctx.BindJSON(&data)

		dates := []time.Time{}
		// loc, _ := time.LoadLocation("Europe/Stockholm")
		for _, record := range data.Records {
			parsed, _ := time.Parse("2006-01-02 15:04:05", record)
			dates = append(dates, parsed)
		}

		tax := calculator.GetTax(model.Car{}, dates)

		ctx.JSON(http.StatusOK, gin.H{"data": tax})
	})

	r.Run()
}
