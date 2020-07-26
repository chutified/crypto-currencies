package models

// Currency defines a basic data model for the cryptocurrency.
type Currency struct {
	Name              string
	Symbol            string
	MarketCapUSD      float64
	Price             float64
	CirculatingSupply float64
	Mineable          bool
	Volume            float64
	ChangeHour        float32
	ChangeDay         float32
	ChangeWeek        float32
}
