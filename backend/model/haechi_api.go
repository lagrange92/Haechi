package model

// PplData : A single population data in Seoul
type PplData struct {
	AreaName      string `json:"areaName"`
	AreaCode      string `json:"areaCode"`
	AreaLatitude  string `json:"areaLatitude"`
	AreaLongitude string `json:"areaLongitude"`
	AreaAvgPpltn  string `json:"areaAvgPpltn"`
}

// CozyPlacesData : A single cozy place data in Seoul
type CozyPlacesData struct {
	AreaName      string `json:"areaName"`
	AreaLatitude  string `json:"areaLatitude"`
	AreaLongitude string `json:"areaLongitude"`
	AreaAvgPpltn  string `json:"areaAvgPpltn"`
}

// ChatPromptData : User sent prompt data to chatbot
type ChatPromptData struct {
	Prompt string `json:"prompt"`
}

// ChatResponseData : Chatbot response data to user
type ChatResponseData struct {
	Chat    string `json:"chat"`
	Suggest string `json:"suggest"`
}
