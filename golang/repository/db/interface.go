package db

import (
	"congestion-calculator/entity"
	"context"
)

type IRepository interface {
	GetCityTaxRule(context.Context, string) (*entity.CityTaxRule, error)
}
