package model

// PplData : A single population data in Seoul
type PplData struct {
	AreaName      string `json:"areaName"`
	AreaCode      string `json:"areaCode"`
	AreaLatitude  string `json:"areaLatitude"`
	AreaLongitude string `json:"areaLongitude"`
	AreaAvgPpltn  string `json:"areaAvgPpltn"`
}
