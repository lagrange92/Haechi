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

// PicnicSpot is a picnic spot for use in picnic planning
type PicnicSpot struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// PicnicRequest is a request for making picnic plan
type PicnicRequest struct {
	Start    PicnicSpot
	LayOvers []PicnicSpot
	Goal     PicnicSpot
}

// PicnicResponse is a response for making picnic plan
type PicnicResponse struct {
	Start    PicnicSpot
	LayOvers []PicnicSpot
	Goal     PicnicSpot
}
