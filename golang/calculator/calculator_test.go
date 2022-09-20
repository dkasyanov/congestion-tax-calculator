package calculator

import (
	"congestion-calculator/model"
	"testing"
	"time"
)

type test struct {
	name           string
	inputVehicle   model.Vehicle
	inputDates     []time.Time
	expectedResult int
}

func TestGetTax(t *testing.T) {
	tests := []test{
		{
			name:           "Car with no records",
			inputVehicle:   model.Car{},
			inputDates:     []time.Time{},
			expectedResult: 0,
		},
		{
			name:         "Car with single record",
			inputVehicle: model.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 12, 15, 7, 0, time.UTC),
			},
			expectedResult: 8,
		},
		{
			name:         "Car with 2 records different days",
			inputVehicle: model.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 12, 15, 7, 0, time.UTC),
				time.Date(2013, 04, 16, 12, 15, 7, 0, time.UTC),
			},
			expectedResult: 8 + 8,
		},
		{
			name:         "Car with 2 records same day, different time",
			inputVehicle: model.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 12, 15, 7, 0, time.UTC),
				time.Date(2013, 04, 15, 15, 15, 7, 0, time.UTC),
			},
			expectedResult: 8 + 13,
		},
		{
			name:         "Car with 2 records same day, in 1 hour",
			inputVehicle: model.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 14, 45, 7, 0, time.UTC),
				time.Date(2013, 04, 15, 15, 05, 7, 0, time.UTC),
			},
			expectedResult: 13,
		},
		{
			name:         "Motorbike with 3 records same day, in 1 hour",
			inputVehicle: model.Motorbike{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 14, 45, 7, 0, time.UTC),
				time.Date(2013, 04, 15, 15, 05, 7, 0, time.UTC),
				time.Date(2013, 04, 15, 15, 07, 7, 0, time.UTC),
			},
			expectedResult: 13,
		},
		{
			name:         "Car with 3 records, 1 on weekend",
			inputVehicle: model.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 4, 14, 45, 7, 0, time.UTC),
				time.Date(2013, 04, 5, 15, 05, 7, 0, time.UTC),
				time.Date(2013, 04, 6, 15, 07, 7, 0, time.UTC),
			},
			expectedResult: 8 + 13 + 0,
		},
	}

	for _, tc := range tests {
		got := GetTax(tc.inputVehicle, tc.inputDates)
		if got != tc.expectedResult {
			t.Fatalf("test: %s: expected: %d, got: %d", tc.name, tc.expectedResult, got)
		}
	}
}
