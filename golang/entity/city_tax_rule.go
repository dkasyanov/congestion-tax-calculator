package entity

type CityTaxRule struct {
	City          string      `bson:"city"`
	DailyMax      int         `bson:"dailyMax"`
	NoTaxWeekdays []string    `bson:"noTaxWeekdays"`
	NoTaxMonth    []string    `bson:"noTaxMonth"`
	NoTaxDates    []string    `bson:"noTaxDates"`
	TaxByTime     []TaxByTime `bson:"taxByTime"`
}

type TaxByTime struct {
	Start  string `bson:"start"`
	End    string `bson:"end"`
	Amount int    `bson:"amount"`
}
