package service

import (
	"congestion-calculator/entity"
	"context"
	"testing"
	"time"
)

type test struct {
	name           string
	inputVehicle   entity.Vehicle
	inputDates     []time.Time
	expectedResult int
}

func TestGetTax(t *testing.T) {
	tests := []test{
		{
			name:           "Car with no records",
			inputVehicle:   entity.Car{},
			inputDates:     []time.Time{},
			expectedResult: 0,
		},
		{
			name:         "Car with single record",
			inputVehicle: entity.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 12, 15, 7, 0, time.UTC),
			},
			expectedResult: 8,
		},
		{
			name:         "Car with 2 records different days",
			inputVehicle: entity.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 12, 15, 7, 0, time.UTC),
				time.Date(2013, 04, 16, 12, 15, 7, 0, time.UTC),
			},
			expectedResult: 8 + 8,
		},
		{
			name:         "Car with 2 records same day, different time",
			inputVehicle: entity.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 12, 15, 7, 0, time.UTC),
				time.Date(2013, 04, 15, 15, 15, 7, 0, time.UTC),
			},
			expectedResult: 8 + 13,
		},
		{
			name:         "Car with 2 records same day, in 1 hour",
			inputVehicle: entity.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 14, 45, 7, 0, time.UTC),
				time.Date(2013, 04, 15, 15, 05, 7, 0, time.UTC),
			},
			expectedResult: 13,
		},
		{
			name:         "Motorcycle with 3 records same day, in 1 hour",
			inputVehicle: entity.Motorcycle{},
			inputDates: []time.Time{
				time.Date(2013, 04, 15, 14, 45, 7, 0, time.UTC),
				time.Date(2013, 04, 15, 15, 05, 7, 0, time.UTC),
				time.Date(2013, 04, 15, 15, 07, 7, 0, time.UTC),
			},
			expectedResult: 13,
		},
		{
			name:         "Car with 3 records, 1 on weekend",
			inputVehicle: entity.Car{},
			inputDates: []time.Time{
				time.Date(2013, 04, 4, 14, 45, 7, 0, time.UTC),
				time.Date(2013, 04, 5, 15, 05, 7, 0, time.UTC),
				time.Date(2013, 04, 6, 15, 07, 7, 0, time.UTC),
			},
			expectedResult: 8 + 13 + 0,
		},
	}

	s := Service{
		conf:                 nil,
		db:                   nil,
		rulesCache:           nil,
		lastRefreshTimestamp: time.Now().UTC(),
	}

	for _, tc := range tests {
		got, err := s.GetTax(context.TODO(), "Gothenburg", tc.inputVehicle, tc.inputDates)
		if got != tc.expectedResult {
			t.Fatalf("test: %s: expected: %d, got: %d", tc.name, tc.expectedResult, got)
		}
		if err != nil {
			t.Fatalf("test: %s: expected: %d, got: error %s", tc.name, tc.expectedResult, err.Error())
		}
	}
}
