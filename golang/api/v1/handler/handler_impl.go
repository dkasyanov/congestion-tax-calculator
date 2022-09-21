package handler

import (
	"congestion-calculator/entity"
	"congestion-calculator/pkg/constants"
	"congestion-calculator/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
			fmt.Println(err.Error())
		}

		dates := []time.Time{}
		for _, record := range data.Records {
			parsed, _ := time.Parse(constants.DateTimeLayout, record)
			dates = append(dates, parsed)
		}

		vehicle := entity.NewVehicle(data.VehicleType)
		tax, err := s.GetTax(ctx, data.City, vehicle, dates)
		if err != nil {
			fmt.Println(err.Error())
		}

		ctx.JSON(http.StatusOK, gin.H{"tax": tax})
	}
}
