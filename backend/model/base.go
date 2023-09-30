package model

// SeoulSpot struct from "./resources/seoul_spot_113.csv"
type SeoulSpot struct {
	Category  string
	ShortCode int
	Code      string
	AreaName  string
	Latitude  float64
	Longitude float64
}

// PplChData data for Ppl Channel
type PplChData struct {
	Spot SeoulSpot
	Data string
}
