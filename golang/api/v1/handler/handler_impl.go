package handler

import (
	"congestion-calculator/entity"
	"congestion-calculator/pkg/constants"
	"congestion-calculator/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	City        string   `json:"city"`
	VehicleType string   `json:"vehicle_type"`
	Records     []string `json:"records"`
}

func postRequestUrl(s service.IService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data RequestBody
		if err := ctx.BindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse input"})
			return
		}

		var dates []time.Time
		for _, record := range data.Records {
			parsed, _ := time.Parse(constants.DateTimeLayout, record)
			dates = append(dates, parsed)
		}

		vehicle := entity.NewVehicle(data.VehicleType)
		tax, err := s.GetTax(ctx, data.City, vehicle, dates)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"message": fmt.Sprintf("Cannot calculate tax. Error: %s", err.Error())},
			)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"tax": tax})
	}
}
