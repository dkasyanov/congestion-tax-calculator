package service

import (
	"congestion-calculator/entity"
	"context"
	"time"
)

type IService interface {
	GetTax(context.Context, string, entity.Vehicle, []time.Time) (int, error)
}
