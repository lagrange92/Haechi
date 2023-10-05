package model

// NVDirResponse : Naver Maps Direction API response data
type NVDirResponse struct {
	Code            int        `json:"code"`
	Message         string     `json:"message"`
	CurrentDateTime string     `json:"currentDateTime"`
	Route           NVDirRoute `json:"route"`
}

// NVDirRoute : Naver Maps Direction API route data
type NVDirRoute struct {
	Trafast []NVDirTrafast `json:"trafast"`
}

// NVDirTrafast : Naver Maps Direction API trafast data
// (trafast: fastest route from start to goal)
type NVDirTrafast struct {
	Summary NVDirSummary   `json:"summary"`
	Path    [][]float64    `json:"path"`
	Section []NVDirSection `json:"section"`
	Guide   []NVDirGuide   `json:"guide"`
}

// NVDirSummary : Naver Maps Direction API summary data
type NVDirSummary struct {
	Start          NVDirLocation `json:"start"`
	Goal           NVDirLocation `json:"goal"`
	Distance       int           `json:"distance"`
	Duration       int           `json:"duration"`
	EtaServiceType int           `json:"etaServiceType"`
	Bbox           [][]float64   `json:"bbox"`
	TollFare       int           `json:"tollFare"`
	TaxiFare       int           `json:"taxiFare"`
	FuelPrice      int           `json:"fuelPrice"`
}

// NVDirLocation : Naver Maps Direction API location data
type NVDirLocation struct {
	Location []float64 `json:"location"`
	Dir      int       `json:"dir,omitempty"`
}

// NVDirSection : Naver Maps Direction API section data
type NVDirSection struct {
	PointIndex int    `json:"pointIndex"`
	PointCount int    `json:"pointCount"`
	Distance   int    `json:"distance"`
	Name       string `json:"name"`
	Congestion int    `json:"congestion"`
	Speed      int    `json:"speed"`
}

// NVDirGuide : Naver Maps Direction API guide data
type NVDirGuide struct {
	PointIndex   int    `json:"pointIndex"`
	Type         int    `json:"type"`
	Instructions string `json:"instructions"`
	Distance     int    `json:"distance"`
	Duration     int    `json:"duration"`
}
