package service

import (
	"congestion-calculator/config"
	"congestion-calculator/entity"
	"congestion-calculator/pkg/constants"
	"congestion-calculator/pkg/utils"
	"congestion-calculator/repository/db"
	"context"
	"fmt"
	"strings"
	"time"
)

type Service struct {
	conf                 *config.Config
	db                   db.IRepository
	rulesCache           map[string]*entity.CityTaxRule
	lastRefreshTimestamp time.Time
}

func New(db db.IRepository, conf *config.Config) *Service {
	return &Service{
		conf:                 conf,
		db:                   db,
		rulesCache:           map[string]*entity.CityTaxRule{},
		lastRefreshTimestamp: time.Time{},
	}
}

func (s *Service) getTaxRules(ctx context.Context, city string) (*entity.CityTaxRule, error) {
	if int(time.Now().UTC().Sub(s.lastRefreshTimestamp).Seconds()) <= s.conf.CacheTTLSeconds {
		rules, ok := s.rulesCache[city]
		if ok {
			return rules, nil
		}
	}

	rules, err := s.db.GetCityTaxRule(ctx, city)
	if err != nil {
		return nil, err
	}

	s.rulesCache[city] = rules
	s.lastRefreshTimestamp = time.Now().UTC()

	return rules, nil
}

func (s *Service) GetTax(ctx context.Context, city string, vehicle entity.Vehicle, dates []time.Time) (int, error) {
	rules, err := s.getTaxRules(ctx, city)
	if err != nil {
		return 0, err
	}

	intervalStart := time.Time{}
	totalFee := 0
	dailyFee := 0
	hourlyFee := 0

	for idx, date := range dates {
		if date.YearDay() != intervalStart.YearDay() {
			// New date, calculate result for the previous one
			dailyFee = dailyFee + hourlyFee
			totalFee = totalFee + utils.Min(dailyFee, rules.DailyTax)
			dailyFee = 0
			hourlyFee = 0
			intervalStart = date
		}
		//fmt.Printf("Date: %s: %d\n", date, dailyFee)

		diffInNanos := date.UnixNano() - intervalStart.UnixNano()
		minutes := diffInNanos / 1000000 / 1000 / 60

		if minutes >= 60 {
			// 60 minutes passed since previous date, adding to daily tax and starting new hour
			dailyFee = dailyFee + hourlyFee
			hourlyFee = 0
			intervalStart = date
		}

		currFee := getTollFee(rules, date, vehicle)
		hourlyFee = utils.Max(hourlyFee, currFee)

		if idx == len(dates)-1 {
			// Last in the list
			dailyFee = dailyFee + hourlyFee
			totalFee = totalFee + utils.Min(dailyFee, rules.DailyTax)
		}
	}

	return totalFee, nil
}

func isTollFreeVehicle(v entity.Vehicle) bool {
	if v == nil {
		return false
	}
	return v.IsTollFree()
}

func getTollFee(rules *entity.CityTaxRule, t time.Time, v entity.Vehicle) int {
	if isTollFreeDate(rules, t) || isTollFreeVehicle(v) {
		return 0
	}

	tString := t.Format(constants.HHMMSSLayout)
	for _, timeRule := range rules.TaxByTime {
		// TODO: add comment
		if tString >= timeRule.Start && tString <= timeRule.End {
			return timeRule.Amount
		}
	}

	return 0
}

func isTollFreeDate(rules *entity.CityTaxRule, date time.Time) bool {
	if utils.SliceContainsString(rules.NoTaxWeekdays, strings.ToUpper(date.Weekday().String())) {
		return true
	}

	if utils.SliceContainsString(rules.NoTaxMonth, strings.ToUpper(date.Month().String())) {
		return true
	}

	for _, d := range rules.NoTaxDates {
		parsedDate, err := time.Parse(constants.YYYYMMDDLayout, d)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		if parsedDate.Year() == date.Year() && parsedDate.YearDay() == date.YearDay() {
			return true
		}
	}

	return false
}
